package batchJobs

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/user/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 增量更新成员
// https://work.weixin.qq.com/api/doc/90000/90135/90980
func (comp *Client) SyncUser(mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callback":  callback,
	}
	_, err := comp.HttpPostJson("cgi-bin/batch/syncuser", options, nil, nil, result)

	return result, err
}

// 全量覆盖成员
// https://work.weixin.qq.com/api/doc/90000/90135/90981
func (comp *Client) ReplaceUser(mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callback":  callback,
	}
	_, err := comp.HttpPostJson("cgi-bin/batch/replaceuser", options, nil, nil, result)

	return result, err
}

// 全量覆盖部门
// https://work.weixin.qq.com/api/doc/90000/90135/90982
func (comp *Client) ReplaceParty(mediaID string, toInvite bool, callback *power.StringMap) (*response.ResponseUserBatchJobs, error) {

	result := &response.ResponseUserBatchJobs{}

	options := &object.HashMap{
		"media_id":  mediaID,
		"to_invite": toInvite,
		"callback":  callback,
	}
	_, err := comp.HttpPostJson("cgi-bin/batch/replaceparty", options, nil, nil, result)

	return result, err
}

// 获取异步任务结果
// https://work.weixin.qq.com/api/doc/90000/90135/90983
func (comp *Client) GetBatchResult(jobID string) (*response.ResponseUserBatchGetResult, error) {

	result := &response.ResponseUserBatchGetResult{}

	params := &object.StringMap{
		"jobid": jobID,
	}
	_, err := comp.HttpGet("cgi-bin/batch/getresult", params, nil, result)

	return result, err
}
