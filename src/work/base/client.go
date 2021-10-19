package base

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/work/base/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90930
func (comp *Client) GetCallbackIp() (*response.ResponseGetCallBackIp, error) {

	result := &response.ResponseGetCallBackIp{}

	_, err := comp.HttpGet("cgi-bin/getcallbackip", nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/92520
func (comp *Client) GetAPIDomainIP() (*response.ResponseGetAPIDomainIP, error) {

	result := &response.ResponseGetAPIDomainIP{}

	_, err := comp.HttpGet("cgi-bin/get_api_domain_ip", nil, nil, result)

	return result, err
}

