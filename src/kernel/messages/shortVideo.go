package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type ShortVideo struct {
	*Video
}

func NewShortVideo(mediaID string, items *object.HashMap) *ShortVideo {
	m := &ShortVideo{
		NewVideo(mediaID, "short_video", items),
	}

	return m
}
