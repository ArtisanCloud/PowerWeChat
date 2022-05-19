package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type MiniProgramPage struct {
	*Message
}

func NewMiniProgramPage(items *object.HashMap) *MiniProgramPage {
	m := &MiniProgramPage{
		NewMessage(items),
	}
	m.Type = "miniProgram_page"

	m.Properties = []string{
		"title",
		"appid",
		"pagepath",
		"thumb_media_id",
	}

	m.SetAttributes(&object.HashMap{
		"required":[]string{
			"thumb_media_id",
			"appid",
			"pagepath",
		},
	})

	return m
}
