package appChat

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat/response"
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

// 创建群聊会话
// https://developer.work.weixin.qq.com/document/path/90245
func (comp *Client) Create(ctx context.Context, options *request.RequestAppChatCreate) (*response.ResponseAppChatCreate, error) {

	result := &response.ResponseAppChatCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/appchat/create", options, nil, nil, result)

	return result, err
}

// 修改群聊会话
// https://developer.work.weixin.qq.com/document/path/90246
func (comp *Client) Update(ctx context.Context, options *request.RequestAppChatUpdate) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/appchat/update", options, nil, nil, result)

	return result, err
}

// 获取群聊会话
// https://developer.work.weixin.qq.com/document/path/90247
func (comp *Client) Get(ctx context.Context, chatID string) (*response.ResponseAppChatGet, error) {

	result := &response.ResponseAppChatGet{}

	params := &object.StringMap{
		"chatid": chatID,
	}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/appchat/get", params, nil, result)

	return result, err
}

// 应用推送消息
// https://developer.work.weixin.qq.com/document/path/90248
func (comp *Client) Send(ctx context.Context, messages *power.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/appchat/send", messages, nil, nil, result)

	return result, err
}
