
.PHONY: build run clean install help

help:
	@echo "Available commands:"
	@echo "  build    - Compile the application"
	@echo "  run      - Run the CLI application"
	@echo "  clean    - Remove built binaries"
	@echo "  install  - Install dependencies"
	@echo "  help     - Show this help message"

build:
	@echo "Building SimpleCLI..."
	go build -o simplecli ./cmd/main.go
	@echo "Build complete! Binary: ./simplecli"

run:
	@echo "Running SimpleCLI..."
	go run ./cmd/main.go $(ARGS)

clean:
	@echo "Cleaning..."
	rm -f simplecli
	@echo "Clean complete!"

install:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed!"
