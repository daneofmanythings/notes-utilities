package zettel

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/daneofmanythings/notes-utilities/internal/static"
	"github.com/daneofmanythings/notes-utilities/internal/utils"
)

func Run() {
	zetName := utils.GetInput("What is the zettel about: ")

	// set up path for zettel
	zetPath := os.Getenv("ZETTEL")
	if zetPath == "" {
		home := os.Getenv("HOME")
		zetPath = filepath.Join(home, "notes", "zettel")
	}

	// adding the filetype and constructing the path
	fileName := genFileName(zetName)
	fullPath := filepath.Join(zetPath, fileName)

	// create file if it does not exist
	createFile(fullPath)

	// echo back the file path
	fmt.Println(fullPath)
}

func createFile(path string) {
	_, err := os.Stat(path)
	if err != nil {
		os.WriteFile(path, static.NewZettel, 0644)
	}
}

func genFileName(s string) string {
	ss := strings.Split(s, " ")
	ss = append(ss, "md")
	return strings.Join(ss, ".")
}
