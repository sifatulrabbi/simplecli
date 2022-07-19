package cli

import (
	"fmt"

	"github.com/sifatulrabbi/simplecli/pkg/constants"
	"github.com/sifatulrabbi/simplecli/pkg/services"
)

func minifyHandler(args []string) {
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
	// If the output file path is given
	outPath := ""
	if len(args) > 5 {
		// Check the syntax
		if len(args) != 7 {
			// If the output path is not specified
			panic(constants.INVALID_ERROR)
		} else {
			outPath = args[6]
		}
	}
	// Find the method and run the minifier
	switch method {
	case "html":
		fmt.Print("Minifying html...\n")
		services.MinifyHTML(path, outPath)
	case "css":
		fmt.Print("Minifying css...\n")
		services.MinifyCSS(path, outPath)
	case "js":
		fmt.Print("Minifying css...\n")
	default:
		fmt.Print(constants.NOT_FOUND_ERROR)
	}
}
