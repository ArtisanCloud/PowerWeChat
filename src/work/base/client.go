package base

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/base/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90930
func (comp *Client) GetCallbackIP() (*response.ResponseGetCallBackIP, error) {

	result := &response.ResponseGetCallBackIP{}

	_, err := comp.HttpGet("cgi-bin/getcallbackip", nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/92520
func (comp *Client) GetAPIDomainIP() (*response.ResponseGetAPIDomainIP, error) {

	result := &response.ResponseGetAPIDomainIP{}

	_, err := comp.HttpGet("cgi-bin/get_api_domain_ip", nil, nil, result)

	return result, err
}
