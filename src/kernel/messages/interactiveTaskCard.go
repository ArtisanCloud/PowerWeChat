package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type InteractiveTaskCard struct {
	*Message
}

func NewInteractiveTaskCard(items *power.HashMap) *InteractiveTaskCard {
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
