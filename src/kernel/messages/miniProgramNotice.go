package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type MiniProgramNotice struct {
	*Message
}

func NewMiniProgramNotice(items *object.HashMap) *MiniProgramNotice {
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
