package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type Music struct {
	*Message
}

func NewMusic(items *power.HashMap) *Music {
	m := &Music{
		NewMessage(items),
	}

	m.Type = "music"

	m.Properties = []string{
		"title",
		"description",
		"url",
		"hq_url",
		"thumb_media_id",
		"format",
	}

	m.JsonAliases = &object.HashMap{
		"musicurl":   "url",
		"hqmusicurl": "hq_url",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *Music) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {

		music := &object.HashMap{
			"Music": &object.HashMap{
				"Title":       msg.Get("title", ""),
				"Description": msg.Get("description", ""),
				"MusicUrl":    msg.Get("url", ""),
				"HQMusicUrl":  msg.Get("hq_url", ""),
			},
		}

		thumbMediaID := msg.Get("thumb_media_id", nil)
		if thumbMediaID != nil {
			musicObject := (*music)["Music"].(*object.HashMap)
			(*musicObject)["ThumbMediaId"] = thumbMediaID
		}

		return music
	}
}
