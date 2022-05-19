package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Location struct {
	*Message
}

func NewLocation(items *object.HashMap) *Location {
	m := &Location{
		NewMessage(items),
	}
	m.Type = "location"

	m.Properties = []string{
		"latitude",
		"longitude",
		"scale",
		"label",
		"precision",
	}

	return m
}
