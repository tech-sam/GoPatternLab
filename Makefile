# Makefile
.PHONY: run list build clean serve

# Variables
BINARY=gopatternlab

# Serve the web application
serve:
	go run cmd/main.go -serve

# Build the binary
build:
	go build -o $(BINARY) cmd/main.go

# Clean build artifacts
clean:
	rm -f $(BINARY)

# This prevents make from trying to process the arguments as targets
%:
	@: