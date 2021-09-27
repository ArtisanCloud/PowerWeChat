package linkedCorp

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/appChat/linkedCorp/response"
)

type Client struct {
	*kernel.BaseClient
}

// 应用推送消息
// https://work.weixin.qq.com/api/doc/90000/90135/90250
func (comp *Client) Send(messages *power.HashMap) (*response.ResponseLinkCorpMessageSend, error) {

	result := &response.ResponseLinkCorpMessageSend{}

	_, err := comp.HttpPostJson("cgi-bin/linkedcorp/message/send", messages, nil, nil, result)

	return result, err
}
