package cli

import (
	"fmt"
	"strings"

	"github.com/sifatulrabbi/simplecli/pkg/constants"
	"github.com/sifatulrabbi/simplecli/pkg/docs"
	"github.com/sifatulrabbi/simplecli/pkg/services"
)

func Cli(args []string) {
	// Check if the user entered any commands or not.
	if len(args) < 1 {
		fmt.Print("No commands found\n")
		return
	}
	// Showing help or mini tutorial
	if args[0] == "-h" || args[0] == "--help" || args[0] == "help" || args[0] == "h" {
		fmt.Print(docs.HelpDoc)
		return
	}
	// Input: `simpleCLI print Hello`
	// Output: Hello>
	if args[0] == "print" {
		printAnyThing(args)
		return
	}
	// Minify HTML, JSON, CSS and JavaScript files
	if args[0] == "minify" {
		handleMinify(args)
	}
}

func printAnyThing(args []string) {
	values := strings.Join(args[1:], " ")
	fmt.Printf(values + "\n")
}

func handleMinify(args []string) {
	// Example: minify -m html -f ./index.html -o ./output.html
	if l := len(args); l < 5 {
		fmt.Print(constants.INVALID_ERROR)
		return
	}

	// All the required vars
	cmd1, cmd2, method, path := args[1], args[3], args[2], args[4]

	// Check for the syntax
	if cmd1 != "-m" && cmd1 != "--method" || cmd2 != "-f" && cmd2 != "--file" {
		fmt.Print(constants.INVALID_ERROR)
		return
	}
	// Find the method and run the minifier
	switch method {
	case "html":
		fmt.Print("Minifying html...\n")
		// If the output file path is given
		if len(args) > 5 {
			// Check the syntax
			if len(args) != 7 {
				fmt.Print(constants.INVALID_ERROR)
			} else {
				oPath := args[6]
				services.MinifyHTML(path, oPath)
			}
		} else {
			// If the output path is not specified
			services.MinifyHTML(path, "")
		}
	case "css":
		fmt.Print("Minifying css...\n")
	case "js":
		fmt.Print("Minifying css...\n")
	default:
		fmt.Print(constants.NOT_FOUND_ERROR)
	}
}