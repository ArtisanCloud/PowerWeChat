package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type Image struct {
	*Media
}

func NewImage(mediaID string, items *power.HashMap) *Image {
	m := &Image{
		NewMedia(mediaID, "image", items),
	}

	return m
}
