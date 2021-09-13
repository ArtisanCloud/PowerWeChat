package messages

import (
	"encoding/json"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
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

func (msg *Raw) TransformForJsonRequest(appends *power.HashMap, withType bool) (data *power.HashMap) {
	data = &power.HashMap{}

	err := json.Unmarshal([]byte(msg.content()), data)
	if err != nil {
		data = &power.HashMap{}
	}
	return data
}

func (msg *Raw) content() string {
	return msg.Get("content", "").(string)
}
