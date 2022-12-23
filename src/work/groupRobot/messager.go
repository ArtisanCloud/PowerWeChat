package groupRobot

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
)

type Messager struct {
	*Client
	message  *messages.Message
	GroupKey string
}

func NewMessager(client *Client) *Messager {
	return &Messager{
		Client: client,
	}
}

func (comp *Messager) Message(message interface{}) (*Messager, error) {

	switch message.(type) {
	case string:
		message = messages.NewText(message.(string))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		message = messages.NewText(fmt.Sprintf("%d", message))
	case float32, float64, complex64, complex128:
		message = messages.NewText(fmt.Sprintf("%f", message))
	case messages.Message:
		obj := message.(messages.Message)
		comp.message = &obj
	case *messages.Message:
		comp.message = message.(*messages.Message)
	default:
		return nil, errors.New(" Invalid uniformMessage. ")
	}

	return comp, nil
}

func (comp *Messager) ToGroup(groupKey string) *Messager {
	comp.GroupKey = groupKey

	return comp
}
