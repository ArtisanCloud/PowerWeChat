package message

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/message/appChat"
	"github.com/ArtisanCloud/power-wechat/src/work/message/externalContact"
	"github.com/ArtisanCloud/power-wechat/src/work/message/linkedCorp"
	"reflect"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *Messager,
	*appChat.Client,
	*externalContact.Client,
	*linkedCorp.Client,
) {
	config := app.GetConfig()

	client := NewClient(app)

	AppChat := appChat.NewClient(app)
	ExternalContact := externalContact.NewClient(app)
	LinkedCorp := linkedCorp.NewClient(app)

	messager := NewMessager(client)

	agentID := config.Get("agent_id", nil)
	switch reflect.TypeOf(agentID).Kind() {
	case reflect.Int:
		messager.OfAgent(agentID.(int))
	}

	return client, messager,
		AppChat,
		ExternalContact,
		LinkedCorp
}
