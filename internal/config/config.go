// Package config internal/config/config.go
package config

import "flag"

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port  string
	Serve bool // Added this field
}

type DBConfig struct {
	Path         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func Parse() *Config {
	cfg := &Config{
		Server: ServerConfig{},
		DB: DBConfig{
			MaxOpenConns: 25,
			MaxIdleConns: 25,
			MaxIdleTime:  "15m",
		},
	}

	// Parse flags
	flag.BoolVar(&cfg.Server.Serve, "serve", false, "Start web server") // Added serve flag
	flag.StringVar(&cfg.Server.Port, "port", "8080", "Port to run server on")
	flag.StringVar(&cfg.DB.Path, "db-path", "./data/patterns.db", "Database file path")
	flag.Parse()

	return cfg
}
