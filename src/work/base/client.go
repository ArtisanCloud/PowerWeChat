package base

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/base/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90930
func (comp *Client) GetCallbackIp() *response.ResponseGetCallBackIp {

	result := &response.ResponseGetCallBackIp{}

	comp.HttpGet("cgi-bin/getcallbackip", nil,nil, result)

	return result
}


// https://open.work.weixin.qq.com/api/doc/90000/90135/92520
func (comp *Client) GetAPIDomainIP() *response.ResponseGetAPIDomainIP {

	result := &response.ResponseGetAPIDomainIP{}

	comp.HttpGet("cgi-bin/get_api_domain_ip", nil,nil, result)

	return result
}

