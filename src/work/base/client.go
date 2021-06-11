package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/work/base/response"
)

type Client struct {
	*kernel.BaseClient
}

func (comp *Client) GetCallbackIp() *response.ResponseGetCallBackIp {

	result := &response.ResponseGetCallBackIp{}

	comp.HttpGet("cgi-bin/getcallbackip", nil, result)

	return result
}


func (comp *Client) GetAPIDomainIP() *response.ResponseGetAPIDomainIP {

	result := &response.ResponseGetAPIDomainIP{}

	comp.HttpGet("cgi-bin/get_api_domain_ip", nil, result)

	return result
}

