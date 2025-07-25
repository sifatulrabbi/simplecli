# SimpleCLI

A Go-based command line interface tool that provides file minification and code snippet generation capabilities.

## Installation

Build from source:
```bash
go build -o simplecli ./cmd/main.go
```

Or run directly:
```bash
go run ./cmd/main.go [args]
```

## Commands

### Help
Display help information:
```bash
simplecli help
simplecli -h
simplecli --help
```

### Print
Simple text printing utility:
```bash
simplecli print "Hello World"
```

### Minify
Minify HTML, CSS, and JavaScript files:
```bash
# Basic usage
simplecli minify -m <method> -f <input-file>

# With output file
simplecli minify -m <method> -f <input-file> -o <output-file>
```

**Supported methods:**
- `html` - Minify HTML files
- `css` - Minify CSS files  
- `js` - Minify JavaScript files

**Examples:**
```bash
# Minify HTML file
simplecli minify -m html -f ./index.html

# Minify CSS with custom output
simplecli minify -m css -f ./styles.css -o ./dist/styles.min.css

# Minify JavaScript
simplecli minify -m js -f ./script.js -o ./dist/script.min.js
```

### Snippets
Generate code snippets for various frameworks:
```bash
simplecli snippets <framework> <component-type> -n <filename>
```

**Supported frameworks and components:**
- `react`: component, provider
- `next`: component, provider  
- `express`: controller, service

**Examples:**
```bash
# Generate React component
simplecli snippets react component -n MyComponent

# Generate Express controller
simplecli snippets express controller -n UserController
```

## Project Structure

```
simplecli/
   cmd/
      main.go              # Entry point
   pkg/
      cli/                 # CLI command handlers
         clli.go         # Main command routing
         minify-handler.go    # Minification logic
         snippet-handler.go   # Snippet generation logic
      constants/           # Error messages and constants
      docs/               # Built-in help documentation
      services/           # Business logic for minification
         minify-html.go
         minify-css.go
         minify-js.go
      utils/              # Helper utilities
   go.mod
```

## Development

**Build:**
```bash
go build -o simplecli ./cmd/main.go
```

**Test:**
```bash
go test ./...
```

**Format code:**
```bash
go fmt ./...
```

**Update dependencies:**
```bash
go mod tidy
```