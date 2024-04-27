package main

import (
	"flag"

	"github.com/daneofmanythings/notes-utilities/internal/date"
	"github.com/daneofmanythings/notes-utilities/internal/title"
	"github.com/daneofmanythings/notes-utilities/internal/zettel"
)

func main() {
	cmdFlag := flag.String("cmd", "", "specify a command to run")
	languageFlag := flag.String("titlelang", "", "specify a language")

	flag.Parse()
	args := flag.Args()

	switch *cmdFlag {
	case titleFlag:
		title.Run(args, *languageFlag)
	case dateFlag:
		// log.Println(*dateFormatFlag)
		date.Run(args)
	case zetFlag:
		zettel.Run(args)
	default:
		displayHelp()
	}
}
