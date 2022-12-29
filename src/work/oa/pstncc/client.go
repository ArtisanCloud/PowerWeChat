package pstncc

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/pstncc/response"
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

// 发起语音电话
// https://developer.work.weixin.qq.com/document/path/91627
func (comp *Client) Call(ctx context.Context, calleeUserID []string) (*response.ResponsePSTNCCCall, error) {

	result := &response.ResponsePSTNCCCall{}

	options := &object.HashMap{
		"callee_userid": calleeUserID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/pstncc/call", options, nil, nil, result)

	return result, err
}

// 获取接听状态
// https://developer.work.weixin.qq.com/document/path/91628
func (comp *Client) GetStates(ctx context.Context, calleeUserID string, callID string) (*response.ResponsePSTNCCGetStates, error) {

	result := &response.ResponsePSTNCCGetStates{}

	options := &object.HashMap{
		"callee_userid": calleeUserID,
		"callid":        callID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/pstncc/getstates", options, nil, nil, result)

	return result, err
}
