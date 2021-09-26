package uniformMessage

import (
	"errors"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/uniformMessage/request"
)

type Client struct {
	*kernel.BaseClient
}

// 下发小程序和公众号统一的服务消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (comp *Client) Send(toUser string, weAppTemplateMsg *request.WeAppTemplateMsg, mpTemplateMsg *request.MPTemplateMsg) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"touser": toUser,
	}

	if weAppTemplateMsg != nil {
		options = object.MergeHashMap(options, &object.HashMap{
			"weapp_template_msg": weAppTemplateMsg,
		})
	} else if mpTemplateMsg != nil {
		options = object.MergeHashMap(options, &object.HashMap{
			"mp_template_msg": mpTemplateMsg,
		})
	} else {
		return nil, errors.New("please given a valid uniformMessage template. ")
	}

	_, err := comp.HttpPostJson("cgi-bin/groupWelcomeTemplate/wxopen/template/uniform_send", options, nil, nil, result)

	return result, err
}
