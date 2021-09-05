package subscribeMessage

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/subscribeMessage/response"
)

type Client struct {
	*kernel.BaseClient
}

// 组合模板并添加至帐号下的个人模板库
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html
func (comp *Client) AddTemplate(tid string, kidList []int, sceneDesc string) (*response.ResponseSubscribeMessageAddTemplate, error) {

	result := &response.ResponseSubscribeMessageAddTemplate{}

	data := &object.HashMap{
		"tid":       tid,
		"kidList":   kidList,
		"sceneDesc": sceneDesc,
	}

	_, err := comp.HttpPostJson("wxaapi/newtmpl/addtemplate", data, nil, nil, result)

	return result, err
}

// 删除帐号下的个人模板
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html
func (comp *Client) DeleteTemplate(priTmplID string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"priTmplId": priTmplID,
	}

	_, err := comp.HttpPostJson("wxaapi/newtmpl/deltemplate", data, nil, nil, result)

	return result, err
}

// 获取小程序账号的类目
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html
func (comp *Client) GetCategory() (*response.ResponseSubscribeMessageGetCategory, error) {

	result := &response.ResponseSubscribeMessageGetCategory{}

	_, err := comp.HttpGet("wxaapi/newtmpl/getcategory", nil, nil, result)

	return result, err
}

// 获取模板标题下的关键词列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html
func (comp *Client) GetPubTemplateKeyWordsByID(tid string) (*response.ResponseSubscribeMessageGetPubTemplateKeyWordsByID, error) {

	result := &response.ResponseSubscribeMessageGetPubTemplateKeyWordsByID{}

	params := &object.StringMap{
		"tid": tid,
	}
	_, err := comp.HttpGet("wxaapi/newtmpl/getpubtemplatekeywords", params, nil, result)

	return result, err
}

// 获取帐号所属类目下的公共模板标题
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html
func (comp *Client) GetPubTemplateTitleList(ids string, start int, limit int) (*response.ResponseSubscribeMessageGetPubTemplateTitleList, error) {

	result := &response.ResponseSubscribeMessageGetPubTemplateTitleList{}

	params := &object.StringMap{
		"ids":   ids,
		"start": fmt.Sprintf("%d", start),
		"limit": fmt.Sprintf("%d", limit),
	}
	_, err := comp.HttpGet("wxaapi/newtmpl/getpubtemplatetitles", params, nil, result)

	return result, err
}

// 获取当前帐号下的个人模板列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func (comp *Client) GetTemplateList() (*response.ResponseSubscribeMessageGetTemplateList, error) {

	result := &response.ResponseSubscribeMessageGetTemplateList{}

	_, err := comp.HttpGet("wxaapi/newtmpl/gettemplate", nil, nil, result)

	return result, err
}

// 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (comp *Client) Send(toUser string, templateID string, page string,
	miniprogramState string, lang string, msgData *object.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"touser":            toUser,
		"template_id":       templateID,
		"page":              page,
		"miniprogram_state": miniprogramState,
		"lang":              lang,
		"data":              msgData,
	}

	_, err := comp.HttpPostJson("cgi-bin/message/subscribe/send", data, nil, nil, result)

	return result, err
}
