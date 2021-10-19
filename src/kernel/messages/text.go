package messages

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Text struct {
	*Message
}

func NewText(content string) *Text {
	m := &Text{
		NewMessage(&object.HashMap{"content": content}),
	}

	m.Type = "text"
	m.Properties = []string{"content"}
	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *Text) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		return &object.HashMap{
			"Content": msg.Get("content", nil),
		}
	}
}
