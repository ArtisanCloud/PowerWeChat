package messages

import (
"github.com/ArtisanCloud/PowerLibs/object"
)

type TextCard struct {
	*Message
}

func NewTextCard(items *object.HashMap) *TextCard {
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
