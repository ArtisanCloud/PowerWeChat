package serviceState

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/serviceState/response"
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

// 获取会话状态
// https://developer.work.weixin.qq.com/document/path/94669
func (comp *Client) Get(ctx context.Context, openKFID string, externalUserID string) (*response.ResponseServiceStateGet, error) {

	result := &response.ResponseServiceStateGet{}

	options := &object.HashMap{
		"open_kfid":       openKFID,
		"external_userid": externalUserID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/service_state/get", options, nil, nil, result)

	return result, err
}

// 变更会话状态
// https://developer.work.weixin.qq.com/document/path/94669
func (comp *Client) Trans(ctx context.Context, openKFID string, externalUserID string, serviceState int, servicerUserID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"open_kfid":       openKFID,
		"external_userid": externalUserID,
		"service_state":   serviceState,
		"servicer_userid": servicerUserID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/service_state/trans", options, nil, nil, result)

	return result, err
}
