package main

import (
	"os"

	"github.com/sifatulrabbi/simplecli/internal/cli"
)

func main() {
	// All the cmd arguments
	var args []string = os.Args[1:]
	cli.RunCli(args)
}
