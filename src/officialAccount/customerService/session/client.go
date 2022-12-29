package session

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/customerService/session/response"
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

// 获取客服会话列表
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) List(ctx context.Context, account string) (*response.ResponseKFSessionList, error) {
	result := &response.ResponseKFSessionList{}

	_, err := comp.BaseClient.HttpGet(ctx, "customservice/kfsession/getsessionlist", &object.StringMap{"kf_account": account}, nil, &result)

	return result, err
}

// 获取未接入会话列表
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Waiting(ctx context.Context) (*response.ResponseKFSessionWaitCaseList, error) {
	result := &response.ResponseKFSessionWaitCaseList{}

	_, err := comp.BaseClient.HttpGet(ctx, "customservice/kfsession/getwaitcase", nil, nil, &result)

	return result, err
}

// 创建会话
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Create(ctx context.Context, account string, openID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &power.HashMap{
		"kf_account": account,
		"openid":     openID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "customservice/kfsession/create", params, nil, nil, &result)

	return result, err
}

// 关闭会话
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Close(ctx context.Context, account string, openID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &power.HashMap{
		"kf_account": account,
		"openid":     openID,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "customservice/kfsession/close", params, nil, nil, &result)

	return result, err
}

// 获取客户会话状态
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Get(ctx context.Context, openID string) (*response.ResponseKFSessionGet, error) {
	result := &response.ResponseKFSessionGet{}

	_, err := comp.BaseClient.HttpGet(ctx, "customservice/kfsession/getsession", &object.StringMap{"openid": openID}, nil, &result)

	return result, err
}
