package updatableMessage

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/updatableMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/updatableMessage/response"
)

type Client struct {
	*kernel.BaseClient
}

// 创建被分享动态消息或私密消息的 activity_id
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html
func (comp *Client) CreateActivityID(unionID string, openID string) (*response.ResponseActiveMessageCreateActiveID, error) {

	result := &response.ResponseActiveMessageCreateActiveID{}

	params := &object.StringMap{
		"unionid": unionID,
		"openid":  openID,
	}

	_, err := comp.HttpGet("cgi-bin/message/wxopen/activityid/create", params, nil, result)

	return result, err
}

// 修改被分享的动态消息。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html
func (comp *Client) SetUpdatableMsg(options *request.RequestSetUpdatableMsg) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.HttpPostJson("cgi-bin/message/wxopen/updatablemsg/send", options, nil, nil, result)

	return result, err
}
