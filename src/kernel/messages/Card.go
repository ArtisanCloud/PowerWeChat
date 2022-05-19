package messages

import "github.com/ArtisanCloud/PowerLibs/object"

type Card struct {
	*Message
}

func NewCard(cardID string) *Card {
	m := &Card{
		NewMessage(&object.HashMap{"card_id": cardID}),
	}

	m.Type = "wxcard"

	m.Properties = []string{"card_id"}

	return m
}
