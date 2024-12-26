// internal/server/server.go
package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tech-sam/GoPatternLab/internal/config"
	"github.com/tech-sam/GoPatternLab/pkg/db"
	"github.com/tech-sam/GoPatternLab/web"
)

func Start(cfg *config.Config) error {
	// Initialize database
	database, err := initDB(cfg.DB)
	if err != nil {
		return fmt.Errorf("database initialization failed: %w", err)
	}
	defer database.Close()

	// Initialize server
	srv, err := initServer(cfg.Server.Port, database)
	if err != nil {
		return fmt.Errorf("server initialization failed: %w", err)
	}

	return run(srv)
}

func initDB(cfg config.DBConfig) (*db.DB, error) {
	return db.New(db.Config{
		Path:         cfg.Path,
		MaxOpenConns: cfg.MaxOpenConns,
		MaxIdleConns: cfg.MaxIdleConns,
		MaxIdleTime:  cfg.MaxIdleTime,
	})
}

func initServer(port string, database *db.DB) (*http.Server, error) {
	handler, err := web.NewHandler(database)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}, nil
}

func run(srv *http.Server) error {
	// Start server
	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("Starting server on http://localhost%s", srv.Addr)
		serverErrors <- srv.ListenAndServe()
	}()

	// Wait for interrupt signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Wait for shutdown or error
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		return gracefulShutdown(srv, sig)
	}
}

func gracefulShutdown(srv *http.Server, sig os.Signal) error {
	log.Printf("Starting shutdown (signal: %v)...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		srv.Close()
		return fmt.Errorf("could not stop server gracefully: %w", err)
	}

	log.Printf("Server stopped gracefully")
	return nil
}
