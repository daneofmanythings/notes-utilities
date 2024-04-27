package static

import (
	_ "embed"
	"io/fs"
)

//go:embed new_zettel.txt
var NewZettel []byte
var NewJournal []byte

var NotePerms fs.FileMode = 0644
