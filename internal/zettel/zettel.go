package zettel

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/daneofmanythings/notes-utilities/internal/static"
)

func Run(zetArgs []string) {
	zetPath := os.Getenv("ZETTEL")
	if zetPath == "" {
		home := os.Getenv("HOME")
		zetPath = filepath.Join(home, "notes", "zettel")
	}

	fileName := genFileName(zetArgs)
	fullPath := filepath.Join(zetPath, fileName)

	createFile(fullPath)

	fmt.Println(fullPath) // the full path is returned to be piped.
}

func genFileName(args []string) string {
	normArgs := normalizeArgs(args)
	normArgs = append(normArgs, "md")
	return strings.Join(normArgs, ".")
}

// Checks if the input has any spaces in it and replaces it with '.'
// When called from the command line, there will be no spaces, and when called fromve
// vim, there will be, but args will be a slice of one. If for some reason, len(args) > 1
// and args[0] has a " ", args[1:] will be ignored.
func normalizeArgs(args []string) []string {
	if strings.Contains(args[0], " ") {
		return strings.Split(args[0], " ")
	}
	return args
}

func createFile(path string) {
	_, err := os.Stat(path)
	if err != nil {
		os.WriteFile(path, static.NewZettel, static.NotePerms)
	}
}
