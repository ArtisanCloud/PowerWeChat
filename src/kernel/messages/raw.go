package messages

import (
	"encoding/json"
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Raw struct {
	*Message
}

func NewRaw(content string) *Raw {
	m := &Raw{
		NewMessage(&object.HashMap{"content": nil}),
	}

	m.TransformForJsonRequest = m.OverrideTransformForJsonRequest

	return m
}

func (msg *Raw) OverrideTransformForJsonRequest(appends *object.HashMap, withType bool) (data *object.HashMap, err error) {

	data = &object.HashMap{}

	err = json.Unmarshal([]byte(msg.content()), data)
	if err != nil {
		data = &object.HashMap{}
	}
	return data, err

}

func (msg *Raw) content() string {
	return msg.Get("content", "").(string)
}
