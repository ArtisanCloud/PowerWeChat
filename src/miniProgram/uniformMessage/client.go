package uniformMessage

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/uniformMessage/request"
)

type Client struct {
	*kernel.BaseClient
}

// 下发小程序和公众号统一的服务消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (comp *Client) Send(options *request.RequestUniformMessageSend) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.HttpPostJson("cgi-bin/message/wxopen/template/uniform_send", options, nil, nil, result)

	return result, err
}
