package base

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/base/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// https://developer.work.weixin.qq.com/document/path/90930
func (comp *Client) GetCallbackIP(ctx context.Context) (*response.ResponseGetCallBackIP, error) {

	result := &response.ResponseGetCallBackIP{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/getcallbackip", nil, nil, result)

	return result, err
}

// https://developer.work.weixin.qq.com/document/path/92520
func (comp *Client) GetAPIDomainIP(ctx context.Context) (*response.ResponseGetAPIDomainIP, error) {

	result := &response.ResponseGetAPIDomainIP{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/get_api_domain_ip", nil, nil, result)

	return result, err
}
