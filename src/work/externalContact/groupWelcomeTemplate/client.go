package groupWelcomeTemplate

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/groupWelcomeTemplate/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/groupWelcomeTemplate/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 发送新客户欢迎语
// https://work.weixin.qq.com/api/doc/90000/90135/92366
func (comp *Client) AddGroupWelcomeTemplate(options *request.RequestGroupWelcomeTemplateAdd) (*response.ResponseAddGroupWelcomeTemplate, error) {

	result := &response.ResponseAddGroupWelcomeTemplate{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/group_welcome_template/add", options, nil, nil, result)

	return result, err
}

// 编辑入群欢迎语素材
// https://work.weixin.qq.com/api/doc/90000/90135/92366
func (comp *Client) EditGroupWelcomeTemplate(options *request.RequestGroupWelcomeTemplateEdit) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/group_welcome_template/edit", options, nil, nil, result)

	return result, err
}

// 获取入群欢迎语素材
// https://work.weixin.qq.com/api/doc/90000/90135/92366
func (comp *Client) GetGroupWelcomeTemplate(templateID string) (*response.ResponseGetGroupWelcomeTemplate, error) {

	result := &response.ResponseGetGroupWelcomeTemplate{}

	options := &object.HashMap{
		"template_id": templateID,
	}
	_, err := comp.HttpPostJson("cgi-bin/externalcontact/group_welcome_template/get", options, nil, nil, result)

	return result, err
}


// 删除入群欢迎语素材
// https://work.weixin.qq.com/api/doc/90000/90135/92366
func (comp *Client) DelGroupWelcomeTemplate(templateID string, agentID int64) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"template_id": templateID,
		"agentid": agentID,
	}
	_, err := comp.HttpPostJson("cgi-bin/externalcontact/group_welcome_template/del", options, nil, nil, result)

	return result, err
}