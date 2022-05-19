package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type InteractiveTaskCard struct {
	*Message
}

func NewInteractiveTaskCard(items *object.HashMap) *InteractiveTaskCard {
	m := &InteractiveTaskCard{
		NewMessage(items),
	}
	m.Type = "interactive_taskcard"

	m.Properties = []string{
		"title",
		"description",
		"url",
		"task_id",
		"btn",
	}

	return m
}
