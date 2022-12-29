package dial

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/dial/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/dial/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取公费电话拨打记录
// https://developer.work.weixin.qq.com/document/path/93662
func (comp *Client) GetDialRecord(ctx context.Context, options *request.RequestDialGetDialRecord) (*response.ResponseDialGetDialRecord, error) {

	result := &response.ResponseDialGetDialRecord{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/dial/get_dial_record", options, nil, nil, result)

	return result, err
}
