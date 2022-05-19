package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type File struct {
	*Message
}

func NewFile(items *object.HashMap) *File {
	m := &File{
		NewMessage(items),
	}
	m.Type = "file"

	return m
}
