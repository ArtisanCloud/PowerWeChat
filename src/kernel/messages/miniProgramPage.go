package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type MiniProgramPage struct {
	*Message
}

func NewMiniProgramPage(items *power.HashMap) *MiniProgramPage {
	m := &MiniProgramPage{
		NewMessage(items),
	}
	m.Type = "miniprogrampage"

	m.Properties = []string{
		"title",
		"appid",
		"pagepath",
		"thumb_media_id",
	}

	m.SetAttribute("required", []string{
		"thumb_media_id",
		"appid",
		"pagepath",
	})

	return m
}
