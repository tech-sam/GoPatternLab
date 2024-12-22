// cmd/main.go
package main

import (
	"github.com/tech-sam/GoPatternLab/internal/config"
	"github.com/tech-sam/GoPatternLab/internal/server"
	"log"
)

func main() {
	// Parse configuration
	cfg := config.Parse()

	// Only start server if serve flag is true
	if cfg.Server.Serve {
		if err := server.Start(cfg); err != nil {
			log.Fatal(err)
		}
		return
	}

	// If serve flag is not set, show usage
	log.Println("Use -serve flag to start the server")
	log.Println("Example: go run cmd/main.go -serve")
}
