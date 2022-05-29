package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
)

type ShortVideo struct {
	*Video
}

func NewShortVideo(mediaID string, items *power.HashMap) *ShortVideo {
	m := &ShortVideo{
		NewVideo(mediaID, items),
	}
	m.Type = "short_video"

	return m
}
