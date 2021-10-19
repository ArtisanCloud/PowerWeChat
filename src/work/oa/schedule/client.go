package schedule

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/powerwechat/src/kernel/response"
	"github.com/ArtisanCloud/powerwechat/src/work/oa/schedule/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 创建日程
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func (comp *Client) Add(schedule *power.HashMap, agentID int) (*response.ResponseScheduleAdd, error) {

	result := &response.ResponseScheduleAdd{}

	options := &object.HashMap{
		"schedule": schedule,
		"agentid":  agentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/schedule/add", options, nil, nil, result)

	return result, err
}

// 更新日程
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func (comp *Client) Update(schedule *power.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"schedule": schedule,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/schedule/update", options, nil, nil, result)

	return result, err
}

// 获取日程详情
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func (comp *Client) Get(scheduleIDList []string) (*response.ResponseScheduleGet, error) {

	result := &response.ResponseScheduleGet{}

	options := &object.HashMap{
		"schedule_id_list": scheduleIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/schedule/get", options, nil, nil, result)

	return result, err
}

// 删除日程
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func (comp *Client) Del(scheduleID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"schedule_id": scheduleID,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/schedule/del", options, nil, nil, result)

	return result, err
}
