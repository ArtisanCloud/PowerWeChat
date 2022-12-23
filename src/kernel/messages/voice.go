package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type Voice struct {
	*Media
}

func NewVoice(mediaID string, item *power.HashMap) *Voice {
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
