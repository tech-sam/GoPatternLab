# Makefile
.PHONY: run list build clean

# Variables
BINARY=blind75

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY) cmd/main.go

# Run with pattern and problem
run:
	$(eval PATTERN := $(wordlist 2,2,$(MAKECMDGOALS)))
	$(eval PROBLEM := $(wordlist 3,3,$(MAKECMDGOALS)))
	@if [ "$(PATTERN)" != "" ] && [ "$(PROBLEM)" != "" ]; then \
		go run cmd/main.go -pattern $(PATTERN) -problem $(PROBLEM); \
	else \
		echo "Usage: make run <pattern> <problem>"; \
		echo "Example: make run twopointers validpalindrome"; \
		exit 1; \
	fi

# List all available patterns and problems
list:
	go run cmd/main.go -list

# Clean build artifacts
clean:
	rm -f $(BINARY)

# This prevents make from trying to process the arguments as targets
%:
	@: