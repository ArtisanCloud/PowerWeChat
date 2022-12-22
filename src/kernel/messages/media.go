package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type Media struct {
	*Message
}

func NewMedia(mediaID string, sType string, items *power.HashMap) *Media {

	data := power.MergeHashMap(&power.HashMap{
		"media_id": mediaID,
	}, items)
	m := &Media{
		NewMessage(data),
	}

	m.SetType(sType)

	m.Properties = []string{
		"media_id",
	}

	m.SetAttribute("required", []string{
		"media_id",
	})

	m.OverrideToXmlArray()

	return m
}

func (msg *Media) GetMediaID() string {

	_ = msg.CheckRequiredAttributes()

	return msg.GetString("media_id", "")
}

// Override ToXmlArray
func (msg *Media) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		return &object.HashMap{
			object.Studly(msg.GetType()): &object.HashMap{
				"MediaId": msg.GetString("media_id", ""),
			},
		}
	}
}
