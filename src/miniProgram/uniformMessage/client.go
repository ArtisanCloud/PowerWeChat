package uniformMessage

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/uniformMessage/request"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 下发小程序和公众号统一的服务消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (comp *Client) Send(ctx context.Context, options *request.RequestUniformMessageSend) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/wxopen/template/uniform_send", options, nil, nil, result)

	return result, err
}
