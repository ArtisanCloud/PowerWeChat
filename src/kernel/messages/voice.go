package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Voice struct {
	*Media
}

func NewVoice(mediaID string, item *object.HashMap) *Voice {
	m := &Voice{
		NewMedia(mediaID, "voice", item),
	}

	m.Type = "voice"

	m.Properties = []string{
		"media_id",
		"recognition",
	}

	return m
}
