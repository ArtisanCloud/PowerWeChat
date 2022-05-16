package autoReply

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/autoReply/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取公众号的自动回复规则
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Getting_Rules_for_Auto_Replies.html
func (comp *Client) Current(mediaID string) (*response.ResponseGettingRulesForAutoReplies, error) {

	//result :=

	//return response.ResponseGettingRulesForAutoReplies, err
	return nil, nil
}
