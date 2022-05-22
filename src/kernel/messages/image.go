package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Image struct {
	*Media
}

func NewImage(mediaID string, items *object.HashMap) *Image {
	m := &Image{
		NewMedia(mediaID, "image", items),
	}

	return m
}
