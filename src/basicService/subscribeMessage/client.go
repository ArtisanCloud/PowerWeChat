package subscribeMessage

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage/request"
	response3 "github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/device/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 组合模板并添加至帐号下的个人模板库
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html
func (comp *Client) AddTemplate(ctx context.Context, tid string, kidList []int, sceneDesc string) (*response3.ResponseSubscribeMessageAddTemplate, error) {

	result := &response3.ResponseSubscribeMessageAddTemplate{}

	data := &object.HashMap{
		"tid":       tid,
		"kidList":   kidList,
		"sceneDesc": sceneDesc,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/newtmpl/addtemplate", data, nil, nil, result)

	return result, err
}

// 删除帐号下的个人模板
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html
func (comp *Client) DeleteTemplate(ctx context.Context, priTmplID string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"priTmplId": priTmplID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/newtmpl/deltemplate", data, nil, nil, result)

	return result, err
}

// 获取小程序账号的类目
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html
func (comp *Client) GetCategory(ctx context.Context) (*response3.ResponseSubscribeMessageGetCategory, error) {

	result := &response3.ResponseSubscribeMessageGetCategory{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/newtmpl/getcategory", nil, nil, result)

	return result, err
}

// 获取模板标题下的关键词列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html
func (comp *Client) GetPubTemplateKeyWordsByID(ctx context.Context, tid string) (*response3.ResponseSubscribeMessageGetPubTemplateKeyWordsByID, error) {

	result := &response3.ResponseSubscribeMessageGetPubTemplateKeyWordsByID{}

	params := &object.StringMap{
		"tid": tid,
	}
	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/newtmpl/getpubtemplatekeywords", params, nil, result)

	return result, err
}

// 获取帐号所属类目下的公共模板标题
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html
func (comp *Client) GetPubTemplateTitleList(ctx context.Context, ids string, start int, limit int) (*response3.ResponseSubscribeMessageGetPubTemplateTitleList, error) {

	result := &response3.ResponseSubscribeMessageGetPubTemplateTitleList{}

	params := &object.StringMap{
		"ids":   ids,
		"start": fmt.Sprintf("%d", start),
		"limit": fmt.Sprintf("%d", limit),
	}
	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/newtmpl/getpubtemplatetitles", params, nil, result)

	return result, err
}

// 获取当前帐号下的个人模板列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func (comp *Client) GetTemplateList(ctx context.Context) (*response3.ResponseSubscribeMessageGetTemplateList, error) {

	result := &response3.ResponseSubscribeMessageGetTemplateList{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/newtmpl/gettemplate", nil, nil, result)

	return result, err
}

// 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (comp *Client) Send(ctx context.Context, data *request.RequestSubscribeMessageSend) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/subscribe/send", data, nil, nil, result)

	return result, err
}

// send发送订阅通知
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#send发送订阅通知
func (comp *Client) BizSend(ctx context.Context, data *request.RequestSubscribeMessageBizSend) (*response.BaseResp, error) {

	result := &response.BaseResp{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/subscribe/bizsend", data, nil, nil, result)

	return result, err
}
