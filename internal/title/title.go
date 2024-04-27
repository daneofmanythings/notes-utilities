package title

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/daneofmanythings/notes-utilities/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func titleify(s string) string {
	t := cases.Title(language.English)
	S := t.String(s)

	return S
}

func Run(args []string, languageFlag string) {
	var s string
	var err error

	if len(args) > 0 {
		s = strings.Join(args, " ")
	}
	if s == "" {
		s, err = utils.ReadLine(os.Stdin)
	}
	if err != nil {
		log.Print(err)
		return
	}

	// TODO: incorporate language flag
	fmt.Println(titleify(s))
}
