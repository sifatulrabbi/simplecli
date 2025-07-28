package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/sifatulrabbi/simplecli/internal/docs"
	sessioncreator "github.com/sifatulrabbi/simplecli/internal/tmux/session_creator"
	"github.com/sifatulrabbi/simplecli/internal/utils"
)

func RunCli(args []string) {
	fmt.Println(len(args), args)
	// Check if the user entered any commands or not.
	if len(args) < 1 {
		fmt.Print("No commands found\n")
		return
	}
	switch args[0] {
	case "-h":
	case "--help":
	case "h":
	case "help":
		// Showing help or mini tutorial
		fmt.Print(docs.HelpDoc)
	case "print":
		handlePrint(args)
	case "minify":
		// Minify HTML, JSON, CSS and JavaScript files
		minifyHandler(args)
	case "snippets":
		// Create snippets handler
		snippetHandler(args)
	case "tmux-project":
	case "tmux-session":
		handleTmuxSession(args[1:])
	default:
		fmt.Println("Error: invalid action.")
	}
}

func handlePrint(args []string) {
	values := strings.Join(args[1:], " ")
	fmt.Printf(values + "\n")
}

func handleTmuxSession(args []string) {
	isTestMode := args[0] == "--test"
	s := utils.TernaryFnCall(isTestMode, sessioncreator.NewTestMode, sessioncreator.New)
	if err := s.CreateNewSession(
		utils.TernaryFnCall(
			isTestMode,
			func() string { return args[1] },
			func() string { return args[0] },
		),
	); err != nil {
		log.Fatalln("Error:", err)
	}
}
