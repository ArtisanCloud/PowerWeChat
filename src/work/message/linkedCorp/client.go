package linkedCorp

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/linkedCorp/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
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

// 应用推送消息
// https://developer.work.weixin.qq.com/document/path/90250
func (comp *Client) Send(ctx context.Context, messages *power.HashMap) (*response.ResponseLinkCorpMessageSend, error) {

	result := &response.ResponseLinkCorpMessageSend{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/linkedcorp/message/send", messages, nil, nil, result)

	return result, err
}
