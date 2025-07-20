# SimpleCLI Makefile

.PHONY: build build-cli build-tui run run-cli run-tui clean install help

# Default target
help:
	@echo "Available commands:"
	@echo "  build      - Compile both CLI and TUI applications"
	@echo "  build-cli  - Compile the CLI application only"
	@echo "  build-tui  - Compile the TUI application only"
	@echo "  run        - Run the CLI application"
	@echo "  run-cli    - Run the CLI application"
	@echo "  run-tui    - Run the TUI application"
	@echo "  clean      - Remove built binaries"
	@echo "  install    - Install dependencies"
	@echo "  help       - Show this help message"

# Build both applications
build: build-cli build-tui

# Build the CLI application
build-cli:
	@echo "Building SimpleCLI..."
	@mkdir -p build
	go build -o build/simplecli ./cmd/old/main.go
	@echo "CLI build complete! Binary: ./build/simplecli"

# Build the TUI application
build-tui:
	@echo "Building SimpleCLI TUI..."
	@mkdir -p build
	go build -o build/simplecli-tui ./cmd/tui/main.go
	@echo "TUI build complete! Binary: ./build/simplecli-tui"

# Run the CLI application (default)
run: run-cli

# Run the CLI application
run-cli:
	go run ./cmd/old/main.go $(ARGS)

# Run the TUI application
run-tui:
	go run ./cmd/tui/main.go $(ARGS)

# Clean built binaries
clean:
	@echo "Cleaning..."
	rm -rf build
	@echo "Clean complete!"

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed!"
