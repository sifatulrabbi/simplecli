# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

SimpleCLI is a Go-based command line interface tool that provides file minification and code snippet generation capabilities. The project follows standard Go project structure with a modular architecture.

## Build and Development Commands

This is a Go project using Go modules. Key commands:

```bash
# Build the project
go build -o simplecli ./cmd/main.go

# Run the project directly
go run ./cmd/main.go [args]

# Test the project (if tests exist)
go test ./...

# Format code
go fmt ./...

# Get dependencies
go mod tidy
```

## Architecture

The codebase follows a clean modular structure:

- **Entry Point**: `cmd/main.go` - Simple main function that passes command line arguments to the CLI handler
- **CLI Logic**: `pkg/cli/` - Contains the main CLI routing and command handlers
- **Services**: `pkg/services/` - Business logic for minification operations (HTML, CSS, JS)
- **Documentation**: `pkg/docs/` - Built-in help documentation
- **Constants**: `pkg/constants/` - Error messages and configuration constants
- **Utilities**: `pkg/utils/` - Helper functions for file operations

## Core Functionality

The CLI supports three main commands:

1. **print** - Simple text printing utility
2. **minify** - File minification for HTML, CSS, and JavaScript
   - Usage: `simplecli minify -m [html|css|js] -f <input-file> [-o <output-file>]`
3. **snippets** - Code snippet generation for React, Next.js, and Express frameworks
   - Usage: `simplecli snippets <framework> <component-type> -n <filename>`

## Key Files for Understanding

- `pkg/cli/clli.go` - Main command routing logic
- `pkg/cli/minify-handler.go` - Handles minification command parsing and execution
- `pkg/cli/snippet-handler.go` - Handles snippet generation (note: has incomplete implementation)
- `pkg/docs/documentation.go` - Contains built-in help text

## Development Notes

- The project uses standard Go conventions with package-based organization
- Error handling uses predefined constants from `pkg/constants/`
- The minify functionality supports HTML, CSS, and JavaScript files
- The snippets feature appears to be partially implemented and may need completion
- No external testing framework is currently in use; use `go test` for any future tests