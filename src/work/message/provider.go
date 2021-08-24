package message

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"reflect"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *Messager) {
	config := app.GetConfig()

	client := &Client{
		kernel.NewBaseClient(&app, nil),
	}
	messager := NewMessager(client)

	agentID := config.Get("agent_id", nil)
	switch reflect.TypeOf(agentID).Kind() {
	case reflect.Int:
		messager.OfAgent(agentID.(int))
	}

	return client, messager
}
