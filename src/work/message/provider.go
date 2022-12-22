package message

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/externalContact"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/linkedCorp"
	"reflect"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *Messager,
	*appChat.Client,
	*externalContact.Client,
	*linkedCorp.Client,
	error,
) {
	config := app.GetConfig()

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	AppChat, err := appChat.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	ExternalContact, err := externalContact.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	LinkedCorp, err := linkedCorp.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	messager := NewMessager(client)

	agentID := config.Get("agent_id", nil)
	switch reflect.TypeOf(agentID).Kind() {
	case reflect.Int:
		messager.OfAgent(agentID.(int))
	}

	return client, messager,
		AppChat,
		ExternalContact,
		LinkedCorp,
		nil
}
