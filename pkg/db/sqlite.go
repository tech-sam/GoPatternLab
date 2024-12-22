// Package db pkg/db/db.go
package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	*sql.DB
}

type Config struct {
	Path         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func New(cfg Config) (*DB, error) {
	// Ensure database directory exists
	if err := os.MkdirAll(filepath.Dir(cfg.Path), 0755); err != nil {
		return nil, fmt.Errorf("create db directory: %w", err)
	}

	db, err := sql.Open("sqlite", cfg.Path)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("parse idle timeout: %w", err)
	}
	db.SetConnMaxIdleTime(duration)

	// Create context with timeout for connection test
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verify database connection
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	return &DB{db}, nil
}

func runMigrations(db *sql.DB) error {
	// First create migrations table
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version TEXT PRIMARY KEY,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return fmt.Errorf("create migrations table: %w", err)
	}

	migrations, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	for _, migration := range migrations {
		if filepath.Ext(migration.Name()) != ".sql" {
			continue
		}

		// Check if migration has been applied
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = ?)",
			migration.Name()).Scan(&exists)
		if err != nil {
			return fmt.Errorf("check migration %s: %w", migration.Name(), err)
		}

		if exists {
			log.Printf("Skipping migration %s (already applied)", migration.Name())
			continue
		}

		// Read and execute migration
		content, err := migrationsFS.ReadFile(filepath.Join("migrations", migration.Name()))
		if err != nil {
			return fmt.Errorf("read migration %s: %w", migration.Name(), err)
		}

		log.Printf("Applying migration: %s", migration.Name())

		// Start transaction for this migration
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("begin transaction for %s: %w", migration.Name(), err)
		}

		// Split migration into separate statements
		statements := strings.Split(string(content), ";")
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			if _, err := tx.Exec(stmt); err != nil {
				tx.Rollback()
				return fmt.Errorf("execute migration %s: %w", migration.Name(), err)
			}
		}

		// Record migration
		if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)",
			migration.Name()); err != nil {
			tx.Rollback()
			return fmt.Errorf("record migration %s: %w", migration.Name(), err)
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit migration %s: %w", migration.Name(), err)
		}
	}

	return nil
}
