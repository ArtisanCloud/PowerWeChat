package externalContact

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/message/externalContact/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 应用推送消息
// https://work.weixin.qq.com/api/doc/90000/90135/90250
func (comp *Client) Send(messages *power.HashMap) (*response.ResponseExternalContactMessageSend, error) {

	result := &response.ResponseExternalContactMessageSend{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/message/send", messages, nil, nil, result)

	return result, err
}
