package message

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/messages"
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
func (comp *Client) Send(messages interface{}) *response.ResponseMessageSend {

	result := &response.ResponseMessageSend{}

	comp.HttpPostJson("cgi-bin/message/send", messages, nil,nil, result)

	return result
}
