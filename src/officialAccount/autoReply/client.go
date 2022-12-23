package autoReply

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/autoReply/response"
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

// 获取公众号的自动回复规则
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Getting_Rules_for_Auto_Replies.html
func (comp *Client) Current() (*response.ResponseGettingRulesForAutoReplies, error) {

	result := &response.ResponseGettingRulesForAutoReplies{}

	_, err := comp.HttpGet("cgi-bin/get_current_autoreply_info", nil, nil, result)

	return result, err

}
