package message

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/response"
	"reflect"
	"strings"
)

type Messager struct {
	message *messages.Message

	To        *power.HashMap
	AgentID   int
	secretive bool
	Client    *Client
}

func NewMessager(client *Client) *Messager {
	m := &Messager{
		Client: client,
	}

	return m
}

func (msg *Messager) Message(message contract.MessageInterface) *Messager {

	msg.message.MessageInterface = message

	return msg

}

func (msg *Messager) OfAgent(agentID int) *Messager {
	msg.AgentID = agentID
	return msg
}

func (msg *Messager) ToUser(userIDs interface{}) *Messager {
	return msg.SetRecipients(userIDs, "touser")
}

func (msg *Messager) ToParty(partyIDs interface{}) *Messager {
	return msg.SetRecipients(partyIDs, "toparty")
}

func (msg *Messager) ToTag(tagIDs interface{}) *Messager {

	return msg.SetRecipients(tagIDs, "totag")
}

func (msg *Messager) Secretive() *Messager {
	msg.secretive = true
	return msg
}
func (msg *Messager) SetRecipients(ids interface{}, key string) *Messager {
	var strIDs string

	switch reflect.TypeOf(ids).Kind() {
	case reflect.Array:
		arrayIDs := ids.([]string)
		strIDs = strings.Join(arrayIDs, "|")

	case reflect.String:
		strIDs = ids.(string)
	default:
	}

	msg.To = &power.HashMap{key: strIDs}

	return msg
}

func (msg *Messager) Send(ctx context.Context, message *messages.Message) (*response.ResponseMessageSend, error) {
	if message != nil {
		msg.Message(message)
	}

	if msg.AgentID <= 0 {
		return nil, errors.New("no agentid specified")
	}

	secretive := 0
	if msg.secretive {
		secretive = 1
	}
	messageToSend, err := msg.message.TransformForJsonRequest(object.MergeHashMap(&object.HashMap{
		"agentid": msg.AgentID,
		"safe":    secretive,
	}, msg.To.ToHashMap()), true)
	if err != nil {
		return nil, err
	}

	msg.secretive = false

	powerMessages, err := power.HashMapToPower(messageToSend)
	if err != nil {
		return nil, err
	}
	return msg.Client.Send(ctx, powerMessages)
}
