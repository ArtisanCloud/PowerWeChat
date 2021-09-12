package message

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/messages"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/message/response"
)

type Client struct {
	*kernel.BaseClient
}

func (comp *Client) Message(message *messages.Message) *Messager {

	m := &Messager{
		Client: comp,
	}
	m.Message(message)

	return m
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90236
func (comp *Client) Send(messages interface{}) (*response.ResponseMessageSend, error) {

	result := &response.ResponseMessageSend{}

	_, err := comp.HttpPostJson("cgi-bin/message/send", messages, nil, nil, result)

	return result, err
}

// Recall 撤回应用消息
// https://open.work.weixin.qq.com/api/doc/90000/90135/94867
func (comp *Client) Recall(msgID string) (*response2.ResponseWork, error) {
	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/message/recall", power.StringMap{
		"msgid": msgID,
	}, nil, nil, result)

	return result, err
}