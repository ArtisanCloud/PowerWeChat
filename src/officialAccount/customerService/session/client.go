package session

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/customerService/session/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取客服会话列表
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) List(account string) (*response.ResponseKFSessionList, error) {
	result := &response.ResponseKFSessionList{}

	_, err := comp.HttpGet("customservice/kfsession/getsessionlist", &object.StringMap{"kf_account": account}, nil, &result)

	return result, err
}

// 获取未接入会话列表
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Waiting() (*response.ResponseKFSessionWaitCaseList, error) {
	result := &response.ResponseKFSessionWaitCaseList{}

	_, err := comp.HttpGet("customservice/kfsession/getwaitcase", nil, nil, &result)

	return result, err
}

// 创建会话
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Create(account string, openID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &power.HashMap{
		"kf_account": account,
		"openid":     openID,
	}
	_, err := comp.HttpPostJson("customservice/kfsession/create", params, nil, nil, &result)

	return result, err
}

// 关闭会话
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Close(account string, openID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	params := &power.HashMap{
		"kf_account": account,
		"openid":     openID,
	}
	_, err := comp.HttpPostJson("customservice/kfsession/close", params, nil, nil, &result)

	return result, err
}

// 获取客户会话状态
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (comp *Client) Get(openID string) (*response.ResponseKFSessionGet, error) {
	result := &response.ResponseKFSessionGet{}

	_, err := comp.HttpGet("customservice/kfsession/getsession", &object.StringMap{"openid": openID}, nil, &result)

	return result, err
}
