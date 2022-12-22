package messages

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type TaskCard struct {
	*Message
}

func NewTaskCard(items *power.HashMap) *TaskCard {
	m := &TaskCard{
		NewMessage(items),
	}
	m.Type = "task_card"

	m.Properties = []string{
		"title",
		"description",
		"url",
		"task_id",
		"btn",
	}

	return m
}
