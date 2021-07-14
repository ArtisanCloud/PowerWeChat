package base

import (
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/base/response"
)

type Client struct {
	*payment.BaseClient
}

//
func (comp *Client) Pay() *response.ResponseGetCallBackIp {

	result := &response.ResponseGetCallBackIp{}

	comp.HttpGet("cgi-bin/getcallbackip", nil,nil, result)

	return result
}

