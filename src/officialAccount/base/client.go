package base

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/officialAccount/base/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/API_Call_Limits.html
func (comp *Client) ClearQuota() (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	config := (*comp.App).GetConfig()

	_, err := comp.HttpPostJson("cgi-bin/clear_quota", nil, &object.StringMap{
		"appid": config.GetString("app_id", ""),
	}, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html#2.%20%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1callback%20IP%E5%9C%B0%E5%9D%80
func (comp *Client) GetCallbackIP() (*response.ResponseGetCallBackIp, error) {

	result := &response.ResponseGetCallBackIp{}

	_, err := comp.HttpGet("cgi-bin/getcallbackip", nil, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Network_Detection.html
func (comp *Client) CheckCallbackURL(action string, checkOperator string) (*response.ResponseGetAPIDomainIP, error) {

	result := &response.ResponseGetAPIDomainIP{}

	options := &object.HashMap{
		"action":         action,
		"check_operator": checkOperator,
	}

	_, err := comp.HttpPostJson("cgi-bin/callback/check", options, nil, nil, result)

	return result, err
}
