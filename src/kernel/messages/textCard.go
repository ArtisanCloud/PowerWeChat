package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type TextCard struct {
	*Message
}

func NewTextCard(items *power.HashMap) *TextCard {
	m := &TextCard{
		NewMessage(items),
	}
	m.Type = "text_card"

	m.Properties = []string{
		"title",
		"description",
		"url",
	}

	return m
}
