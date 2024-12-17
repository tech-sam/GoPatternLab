// pkg/db/sqlite.go
package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	*sql.DB
}

func New(dbPath string) (*DB, error) {
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("create db directory: %w", err)
	}

	// Open SQLite database
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	return &DB{db}, nil
}

func runMigrations(db *sql.DB) error {
	migrations, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	for _, migration := range migrations {
		if filepath.Ext(migration.Name()) != ".sql" {
			continue
		}

		content, err := migrationsFS.ReadFile(filepath.Join("migrations", migration.Name()))
		if err != nil {
			return fmt.Errorf("read migration %s: %w", migration.Name(), err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("execute migration %s: %w", migration.Name(), err)
		}
	}

	return nil
}
