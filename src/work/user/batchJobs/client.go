package batchJobs

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/response"
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

// 增量更新成员
// https://developer.work.weixin.qq.com/document/path/90980
func (comp *Client) SyncUser(ctx context.Context, mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callbacks": callback,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/batch/syncuser", options, nil, nil, result)

	return result, err
}

// 全量覆盖成员
// https://developer.work.weixin.qq.com/document/path/90981
func (comp *Client) ReplaceUser(ctx context.Context, mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callbacks": callback,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/batch/replaceuser", options, nil, nil, result)

	return result, err
}

// 全量覆盖部门
// https://developer.work.weixin.qq.com/document/path/90982
func (comp *Client) ReplaceParty(ctx context.Context, mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callbacks": callback,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/batch/replaceparty", options, nil, nil, result)

	return result, err
}

// 获取异步任务结果
// https://developer.work.weixin.qq.com/document/path/90983
func (comp *Client) GetBatchResult(ctx context.Context, jobID string) (*response.ResponseUserBatchGetResult, error) {

	result := &response.ResponseUserBatchGetResult{}

	params := &object.StringMap{
		"jobid": jobID,
	}
	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/batch/getresult", params, nil, result)

	return result, err
}
