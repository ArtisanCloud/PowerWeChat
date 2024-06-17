package accountService

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/response"
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

// 添加客服帐号
// https://developer.work.weixin.qq.com/document/path/94648
func (comp *Client) Add(ctx context.Context, name string, mediaID string) (*response.ResponseAccountAdd, error) {

	result := &response.ResponseAccountAdd{}

	options := &object.HashMap{
		"name":     name,
		"media_id": mediaID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/account/add", options, nil, nil, result)

	return result, err
}

// 删除客服帐号
// https://developer.work.weixin.qq.com/document/path/94663
func (comp *Client) Del(ctx context.Context, openKFID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"open_kfid": openKFID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/account/del", options, nil, nil, result)

	return result, err
}

// 修改客服帐号
// https://developer.work.weixin.qq.com/document/path/94664
func (comp *Client) Update(ctx context.Context, options *request.RequestAccountUpdate) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/account/update", options, nil, nil, result)

	return result, err
}

// 获取客服帐号列表
// https://developer.work.weixin.qq.com/document/path/94661
func (comp *Client) List(ctx context.Context) (*response.ResponseAccountList, error) {

	result := &response.ResponseAccountList{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/account/list", nil, nil, nil, result)

	return result, err
}

// 获取客服帐号链接
// https://developer.work.weixin.qq.com/document/path/94665
func (comp *Client) AddContactWay(ctx context.Context, openKFID string, scene string) (*response.ResponseAccountServiceAddContactWay, error) {

	result := &response.ResponseAccountServiceAddContactWay{}

	var options *object.HashMap

	if scene == "" {
		options = &object.HashMap{
			"open_kfid": openKFID,
		}
	} else {
		options = &object.HashMap{
			"open_kfid": openKFID,
			"scene":     scene,
		}
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/kf/add_contact_way", options, nil, nil, result)

	return result, err
}
