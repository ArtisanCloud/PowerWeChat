package templateMessage

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/templateMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/templateMessage/response"
)

const API_TEMPLATE_MESSAGE_SEND string = "cgi-bin/message/template/send"

type Client struct {
	*kernel.BaseClient

	Message  *object.HashMap
	Required []string
}

// 设置所属行业
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (comp *Client) SetIndustry(industryOne string, industryTwo string, optional *power.HashMap) (*response.ResponseTemplateIndustry, error) {

	result := &response.ResponseTemplateIndustry{}
	params := &object.HashMap{
		"industry_id1": industryOne,
		"industry_id2": industryTwo,
	}
	_, err := comp.HttpPostJson("cgi-bin/template/api_set_industry", params, nil, nil, result)

	return result, err

}

// 获取设置的行业信息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (comp *Client) GetIndustry() (*response.ResponseTemplateIndustry, error) {

	result := &response.ResponseTemplateIndustry{}

	_, err := comp.HttpPostJson("cgi-bin/template/get_industry", nil, nil, nil, result)

	return result, err

}

// 获得模板ID
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (comp *Client) AddTemplate(industryOne string) (*response.ResponseTemplate, error) {

	result := &response.ResponseTemplate{}

	params := &object.HashMap{
		"template_id_short": industryOne,
	}
	_, err := comp.HttpPostJson("cgi-bin/template/api_add_template", params, nil, nil, result)

	return result, err

}

// 获取模板列表
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (comp *Client) GetPrivateTemplates() (*response.ResponseTemplateGetPrivate, error) {

	result := &response.ResponseTemplateGetPrivate{}

	_, err := comp.HttpPostJson("cgi-bin/template/get_all_private_template", nil, nil, nil, result)

	return result, err

}

// 删除模板
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (comp *Client) DeletePrivateTemplate(templateID string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"template_id": templateID,
	}
	_, err := comp.HttpPostJson("cgi-bin/template/del_private_template", params, nil, nil, result)

	return result, err

}

func (comp *Client) Send(message *request.RequestTemlateMessage) (*response.ResponseTemplateSend, error) {

	result := &response.ResponseTemplateSend{}

	_, err := comp.HttpPostJson(API_TEMPLATE_MESSAGE_SEND, message, nil, nil, result)

	return result, err
}

func (comp *Client) SendSubscription(message *request.RequestTemlateMessageSubscribe) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.HttpPostJson("cgi-bin/message/template/subscribe", message, nil, nil, result)

	return result, err
}
