package main

import (
	"flag"
	"fmt"
	"github.com/tech-sam/GoPatternLab/patterns"
	"github.com/tech-sam/GoPatternLab/pkg/db"
	"log"
	"path/filepath"
)

func main() {
	pattern := flag.String("pattern", "", "Pattern name (e.g., twopointers)")
	problem := flag.String("problem", "", "Problem name (e.g., validpalindrome)")
	list := flag.Bool("list", false, "List all available patterns")
	flag.Parse()

	// Initialize database
	dbPath := filepath.Join(".", "data", "patterns.db")
	database, err := db.New(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	registry := patterns.NewRegistry()

	if *list {
		fmt.Println("Available Patterns and Problems:")
		fmt.Println("================================")
		for pattern, problems := range registry.ListPatterns() {
			fmt.Printf("\n🔹 %s:\n", pattern)
			for _, problem := range problems {
				fmt.Printf("  • %s\n", problem)
			}
		}
		return
	}

	if *pattern == "" || *problem == "" {
		fmt.Println("Usage: go run cmd/main.go -pattern <pattern> -problem <problem>")
		fmt.Println("Or: go run cmd/main.go -list")
		return
	}

	if err := registry.RunProblem(*pattern, *problem); err != nil {
		log.Fatal(err)
	}
}
