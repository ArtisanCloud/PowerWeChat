package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Link struct {
	*Message
}

func NewLink(items *object.HashMap) *Link {
	m := &Link{
		NewMessage(items),
	}
	m.Type = "link"

	m.Properties = []string{
		"title",
		"description",
		"url",
	}

	return m
}
