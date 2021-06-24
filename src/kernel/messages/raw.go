package messages

import (
	"encoding/json"
	"github.com/ArtisanCloud/go-libs/object"
)

type Raw struct {
	*Message
}

func NewRaw(content string) *Raw {
	m := &Raw{
		NewMessage(&object.HashMap{"content": nil}),
	}

	return m
}

func (msg *Raw) TransformForJsonRequest(appends *object.HashMap, withType bool) (data *object.HashMap) {
	data = &object.HashMap{}

	err := json.Unmarshal([]byte(msg.content()), data)
	if err != nil {
		data = &object.HashMap{}
	}
	return data
}

func (msg *Raw) content() string {
	return msg.Get("content", "").(string)
}
