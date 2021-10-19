package externalContact

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/work/message/externalContact/response"
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
