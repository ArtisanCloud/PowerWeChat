package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type MiniProgramNotice struct {
	*Message
}

func NewMiniProgramNotice(items *power.HashMap) *MiniProgramNotice {
	m := &MiniProgramNotice{
		NewMessage(items),
	}
	m.Type = "miniProgram_notice"

	m.Properties = []string{
		"appid",
		"title",
	}

	return m
}
