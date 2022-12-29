package schedule

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/schedule/response"
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

// 创建日程
// https://developer.work.weixin.qq.com/document/path/93648
func (comp *Client) Add(ctx context.Context, schedule *power.HashMap, agentID int) (*response.ResponseScheduleAdd, error) {

	result := &response.ResponseScheduleAdd{}

	options := &object.HashMap{
		"schedule": schedule,
		"agentid":  agentID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/schedule/add", options, nil, nil, result)

	return result, err
}

// 更新日程
// https://developer.work.weixin.qq.com/document/path/93648
func (comp *Client) Update(ctx context.Context, schedule *power.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"schedule": schedule,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/schedule/update", options, nil, nil, result)

	return result, err
}

// 获取日程详情
// https://developer.work.weixin.qq.com/document/path/93648
func (comp *Client) Get(ctx context.Context, scheduleIDList []string) (*response.ResponseScheduleGet, error) {

	result := &response.ResponseScheduleGet{}

	options := &object.HashMap{
		"schedule_id_list": scheduleIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/schedule/get", options, nil, nil, result)

	return result, err
}

// 删除日程
// https://developer.work.weixin.qq.com/document/path/93648
func (comp *Client) Del(ctx context.Context, scheduleID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"schedule_id": scheduleID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/schedule/del", options, nil, nil, result)

	return result, err
}
