package servicer

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/servicer/response"
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

// 添加接待人员
// https://developer.work.weixin.qq.com/document/path/94646
func (comp *Client) Add(ctx context.Context, openKFID string, userIDList []string) (*response.ResponseServicerAdd, error) {

	result := &response.ResponseServicerAdd{}

	options := &object.HashMap{
		"open_kfid":   openKFID,
		"userid_list": userIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/servicer/add", options, nil, nil, result)

	return result, err
}

// 删除接待人员
// https://developer.work.weixin.qq.com/document/path/94647
func (comp *Client) Del(ctx context.Context, openKFID string, userIDList []string) (*response.ResponseServicerDel, error) {

	result := &response.ResponseServicerDel{}

	options := &object.HashMap{
		"open_kfid":   openKFID,
		"userid_list": userIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/servicer/del", options, nil, nil, result)

	return result, err
}

// 获取接待人员列表
// https://developer.work.weixin.qq.com/document/path/94645
func (comp *Client) List(ctx context.Context, openKFID string) (*response.ResponseServicerList, error) {

	result := &response.ResponseServicerList{}

	params := &object.StringMap{
		"open_kfid": openKFID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/servicer/list", nil, params, nil, result)

	return result, err
}
