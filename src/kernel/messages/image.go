package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Image struct {
	*Message
}

func NewImage(items *object.HashMap) *Image {
	m := &Image{
		NewMessage(items),
	}
	m.Type = "image"

	return m
}
