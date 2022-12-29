package base

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/base/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/API_Call_Limits.html
func (comp *Client) ClearQuota(ctx *context.Context) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	config := (*comp.BaseClient.App).GetConfig()

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/clear_quota", &object.HashMap{
		"appid": config.GetString("app_id", ""),
	}, nil, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html#2.%20%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1callback%20IP%E5%9C%B0%E5%9D%80
func (comp *Client) GetCallbackIP(ctx *context.Context) (*response.ResponseGetCallBackIP, error) {

	result := &response.ResponseGetCallBackIP{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/getcallbackip", nil, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Network_Detection.html
func (comp *Client) CheckCallbackURL(ctx *context.Context, action string, checkOperator string) (*response.ResponseGetAPIDomainIP, error) {

	result := &response.ResponseGetAPIDomainIP{}

	options := &object.HashMap{
		"action":         action,
		"check_operator": checkOperator,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/callbacks/check", options, nil, nil, result)

	return result, err
}
