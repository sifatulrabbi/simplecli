package main

import (
	"os"

	"github.com/sifatulrabbi/simplecli/pkg/cli"
)

func main() {
	// All the cmd arguments
	var args []string = os.Args[1:]
	cli.Cli(args)
}
