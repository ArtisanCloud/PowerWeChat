package calendar

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/calendar/response"
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

// 创建日历
// https://developer.work.weixin.qq.com/document/path/93647
func (comp *Client) Add(ctx context.Context, calendar *power.HashMap, agentID int) (*response.ResponseCalendarAdd, error) {

	result := &response.ResponseCalendarAdd{}

	options := &object.HashMap{
		"calendar": calendar,
		"agentid":  agentID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/calendar/add", options, nil, nil, result)

	return result, err
}

// 更新日历
// https://developer.work.weixin.qq.com/document/path/93647
func (comp *Client) Update(ctx context.Context, calendar *power.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"calendar": calendar,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/calendar/update", options, nil, nil, result)

	return result, err
}

// 获取日历详情
// https://developer.work.weixin.qq.com/document/path/93647
func (comp *Client) Get(ctx context.Context, calIDList []string) (*response.ResponseCalendarGet, error) {

	result := &response.ResponseCalendarGet{}

	options := &object.HashMap{
		"cal_id_list": calIDList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/calendar/get", options, nil, nil, result)

	return result, err
}

// 删除日历
// https://developer.work.weixin.qq.com/document/path/93647
func (comp *Client) Del(ctx context.Context, calID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"cal_id": calID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/calendar/del", options, nil, nil, result)

	return result, err
}
