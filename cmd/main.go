package main

import (
	"flag"
	"fmt"
	"github.com/tech-sam/grokking-patterns-go-blind-75/patterns"
	"log"
)

func main() {
	pattern := flag.String("pattern", "", "Pattern name (e.g., twopointers)")
	problem := flag.String("problem", "", "Problem name (e.g., validpalindrome)")
	list := flag.Bool("list", false, "List all available patterns")
	flag.Parse()

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
