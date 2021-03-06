package externalContact

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/message/externalContact/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 发送「学校通知」
// https://developer.work.weixin.qq.com/document/path/92291
func (comp *Client) Send(messages *power.HashMap) (*response.ResponseExternalContactMessageSend, error) {

	result := &response.ResponseExternalContactMessageSend{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/message/send", messages, nil, nil, result)

	return result, err
}
