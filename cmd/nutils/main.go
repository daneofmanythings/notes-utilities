package main

import (
	"flag"

	"github.com/daneofmanythings/notes-utilities/internal/title"
)

func main() {
	cmdFlag := flag.String("cmd", "", "specify a command to run")

	flag.Parse()
	args := flag.Args()

	switch *cmdFlag {
	case titleFlag:
		title.Run(args)
	default:
		displayHelp()
	}
}
