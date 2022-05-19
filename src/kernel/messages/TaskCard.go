package messages

import (
"github.com/ArtisanCloud/PowerLibs/object"
)

type TaskCard struct {
	*Message
}

func NewTaskCard(items *object.HashMap) *TaskCard {
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
