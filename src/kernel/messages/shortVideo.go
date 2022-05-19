package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type ShortVideo struct {
	*Message
}

func NewShortVideo(items *object.HashMap) *Music {
	m := &Music{
		NewMessage(items),
	}

	m.Type = "short_video"

	return m
}
