package shakeAround

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/response"
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

// 申请开通摇一摇周边功能
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Apply/Application_for_opening_function.html
func (comp *Client) Register(ctx *context.Context, data *request.RequestShakeAroundAccountRegister) (*response.ResponseShakeAroundAccountRegister, error) {

	result := &response.ResponseShakeAroundAccountRegister{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/account/register", data, nil, nil, result)

	return result, err
}

// 查询审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Apply/Application_for_opening_function.html
func (comp *Client) Status(ctx *context.Context) (*response.ResponseShakeAroundAccountRegister, error) {

	result := &response.ResponseShakeAroundAccountRegister{}

	_, err := comp.BaseClient.HttpGet(ctx, "shakearound/account/auditstatus", nil, nil, result)

	return result, err
}

// 获取设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Getting_Device_and_User_Information.html
func (comp *Client) User(ctx *context.Context, data *request.RequestShakeAroundUser) (*response.ResponseShakeAroundUser, error) {

	result := &response.ResponseShakeAroundUser{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/user/getshakeinfo", data, nil, nil, result)

	return result, err
}

func (comp *Client) UserWithPoi(ctx *context.Context, ticket string) (*response.ResponseShakeAroundUser, error) {

	data := &request.RequestShakeAroundUser{
		Ticket:  ticket,
		NeedPoi: 1,
	}

	return comp.User(ctx, data)
}
