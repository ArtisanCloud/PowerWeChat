package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type Client struct {
	*kernel.BaseClient
}

func (comp *Client) GetCallbackIp() *ResponseGetCallBackIp {

	result := &ResponseGetCallBackIp{}

	comp.HttpGet("cgi-bin/getcallbackip", nil, result)

	return result
}


func (comp *Client) GetAPIDomainIP() *ResponseGetAPIDomainIP {

	result := &ResponseGetAPIDomainIP{}

	comp.HttpGet("cgi-bin/get_api_domain_ip", nil, result)

	return result
}

