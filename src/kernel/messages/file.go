package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type File struct {
	*Media
}

func NewFile(mediaID string, items *object.HashMap) *File {
	m := &File{
		NewMedia(mediaID, "file", items),
	}

	return m
}
