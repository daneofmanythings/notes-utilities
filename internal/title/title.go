package title

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/daneofmanythings/notes-utilities/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func titleify(as string, b *bytes.Buffer) error {
	t := cases.Title(language.English)
	_, err := b.WriteString(t.String(as))
	return err // always nil. pretending it isn't
}

func Titleify(s string) {
	var b bytes.Buffer

	err := titleify(s, &b)
	if err != nil { // if there is an error, just print the input back out
		fmt.Println(s)
	}

	fmt.Println(b.String())
}

func Run(args []string, languageFlag string) {
	var s string
	var err error

	// TODO: add language flag

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

	Titleify(s)
}
