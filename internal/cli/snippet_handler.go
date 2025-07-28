package cli

import (
	"fmt"

	"github.com/sifatulrabbi/simplecli/internal/constants"
)

// Example: simpleCLI snippets react component -n filename
// Indexes:				0		1		2		3	4

// All the supported frameworks and their components.
var frameworks = map[string]map[string]string{
	"react":   {"component": "React component", "provider": "React provider"},
	"next":    {"component": "Next component", "provider": "React provider"},
	"express": {"controller": "Express controller", "service": "Express service"},
}

// Handle the snippets command.
func snippetHandler(args []string) {
	var (
		inFrame  string            // Input framework name.
		compType string            // Input component type (name).
		filename string            // Input filename (not a file path. Only the name without any extension).
		selFrame map[string]string // Selected framework's map.
	)
	if n := args[3]; n != "-n" && n != "--name" { // If the syntax is correct and a filename is provided.
		fmt.Print(constants.INCORRECT_ARGS)
		return
	}
	filename = args[4]             // Append the file name
	for i, v := range frameworks { // Find the framework and append to selFrame variable.
		if i == inFrame {
			selFrame = v
			break
		}
	}
	if selFrame == nil { // If the selected framework data is nil/not found then stop the process and show an error message.
		var frmList []string
		for n := range frameworks {
			frmList = append(frmList, n)
		}
		fmt.Printf("Unable to find your desired framework. Supported frameworks\n%v\n", frmList)
		return
	}
	if v := selFrame[compType]; len(v) < 1 { // If the selected component name is not in the framework's data then show an error.
		fmt.Println("Unable to find your desired component.")
		return
	}

	fmt.Printf("Selected framework: %v,\ncomponent: %v,\nfilename: %v\n", inFrame, selFrame[compType], filename)
}
