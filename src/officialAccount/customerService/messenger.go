package customerService

import (
	"errors"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/messages"
)

type Messenger struct {
	Message *contract.MessageInterface

	To      string
	Account string

	Client *Client
}

func NewMessenger(client *Client) *Messenger {

	return &Messenger{
		Client: client,
	}
}

func (comp *Messenger) SetMessage(message *contract.MessageInterface) *Messenger {
	comp.Message = message

	return comp
}

func (comp *Messenger) SetBy(account string) *Messenger {
	comp.Account = account

	return comp
}

func (comp *Messenger) From(account string) *Messenger {
	return comp.SetTo(account)
}

func (comp *Messenger) SetTo(openID string) *Messenger {
	comp.To = openID
	return comp
}

func (comp *Messenger) Send() (result interface{}, err error) {

	if comp.Message == nil {
		return result, errors.New("no message to send")
	}

	switch (*comp.Message).(type) {
	case *messages.Raw:
		content := (*comp.Message).(*messages.Raw).Get("content", nil)
		json, err := object.StructToJson(content)
		if err != nil {
			return result, err
		}

		result, err = comp.Client.Send(json)

		break

	default:

		prepends := &object.HashMap{
			"touser": comp.To,
		}
		if comp.Account != "" {
			(*prepends)["customservice"] = &object.HashMap{
				"kf_account": comp.Account,
			}
		}
		json, err := (*comp.Message).TransformForJsonRequest(prepends, true)
		if err != nil {
			return result, err
		}
		result, err = comp.Client.Send(json)
	}



	return result, err
}
