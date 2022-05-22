package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Video struct {
	*Media
}

func NewVideo(mediaID string, sType string, item *object.HashMap) *Video {
	if sType==""{
		sType = "video"
	}
	m := &Video{
		NewMedia(mediaID, sType, item),
	}

	m.Properties = []string{
		"title",
		"description",
		"media_id",
		"thumb_media_id",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *Video) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {

		return &object.HashMap{
			"Video": &object.HashMap{
				"MediaId":     msg.GetString("media_id", ""),
				"Title":       msg.GetString("title", ""),
				"Description": msg.GetString("description", ""),
			},
		}

	}
}
