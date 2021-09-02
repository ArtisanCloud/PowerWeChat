package uniformMessage

import (
	"errors"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/uniformMessage/request"
	"github.com/ArtisanCloud/power-wechat/src/work/message/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func (comp *Client) Send(toUser string, weAppTemplateMsg *request.WeAppTemplateMsg, mpTemplateMsg *request.MPTemplateMsg) (*response.ResponseMessageSend, error) {

	result := &response.ResponseMessageSend{}

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

	_, err := comp.HttpPostJson("cgi-bin/uniformMessage/wxopen/template/uniform_send", options, nil, nil, result)

	return result, err
}
