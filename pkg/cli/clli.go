package cli

import (
	"fmt"
	"strings"

	"github.com/sifatulrabbi/simplecli/pkg/docs"
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
		minifyHandler(args)
	}
}

func printAnyThing(args []string) {
	values := strings.Join(args[1:], " ")
	fmt.Printf(values + "\n")
}
