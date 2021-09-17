package base

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/base/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/API_Call_Limits.html
func (comp *Client) ClearQuota() (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	config := (*comp.App).GetConfig()

	_, err := comp.HttpPostJson("cgi-bin/clear_quota", nil, &object.StringMap{
		"appid": config.GetString("app_id", ""),
	}, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90930
func (comp *Client) getValidIPs() (*response.ResponseGetCallBackIp, error) {

	result := &response.ResponseGetCallBackIp{}

	_, err := comp.HttpGet("cgi-bin/getcallbackip", nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/92520
func (comp *Client) CheckCallbackURL() (*response.ResponseGetAPIDomainIP, error) {

	result := &response.ResponseGetAPIDomainIP{}

	_, err := comp.HttpGet("cgi-bin/callback/check", nil, nil, result)

	return result, err
}
