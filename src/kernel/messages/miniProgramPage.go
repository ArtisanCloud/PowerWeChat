package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
)

type MiniProgramPage struct {
	*Message
}

func NewMiniProgramPage(items *power.HashMap) *MiniProgramPage {
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

	m.SetAttribute("required", []string{
		"thumb_media_id",
		"appid",
		"pagepath",
	})

	return m
}
