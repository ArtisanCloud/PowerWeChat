package groupWelcomeTemplate

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupWelcomeTemplate/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupWelcomeTemplate/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 发送新客户欢迎语
// https://developer.work.weixin.qq.com/document/path/92366
func (comp *Client) AddGroupWelcomeTemplate(ctx context.Context, options *request.RequestGroupWelcomeTemplateAdd) (*response.ResponseAddGroupWelcomeTemplate, error) {

	result := &response.ResponseAddGroupWelcomeTemplate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/group_welcome_template/add", options, nil, nil, result)

	return result, err
}

// 编辑入群欢迎语素材
// https://developer.work.weixin.qq.com/document/path/92366#编辑入群欢迎语素材
func (comp *Client) EditGroupWelcomeTemplate(ctx context.Context, options *request.RequestGroupWelcomeTemplateEdit) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/group_welcome_template/edit", options, nil, nil, result)

	return result, err
}

// 获取入群欢迎语素材
// https://developer.work.weixin.qq.com/document/path/92366#获取入群欢迎语素材
func (comp *Client) GetGroupWelcomeTemplate(ctx context.Context, templateID string) (*response.ResponseGetGroupWelcomeTemplate, error) {

	result := &response.ResponseGetGroupWelcomeTemplate{}

	options := &object.HashMap{
		"template_id": templateID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/group_welcome_template/get", options, nil, nil, result)

	return result, err
}

// 删除入群欢迎语素材
// https://developer.work.weixin.qq.com/document/path/92366#删除入群欢迎语素材
func (comp *Client) DelGroupWelcomeTemplate(ctx context.Context, templateID string, agentID int64) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"template_id": templateID,
		"agentid":     agentID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/group_welcome_template/del", options, nil, nil, result)

	return result, err
}
