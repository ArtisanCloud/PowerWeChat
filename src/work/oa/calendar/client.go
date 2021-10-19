package calendar

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/powerwechat/src/kernel/response"
	"github.com/ArtisanCloud/powerwechat/src/work/oa/calendar/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 创建日历
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func (comp *Client) Add(calendar *power.HashMap, agentID int) (*response.ResponseCalendarAdd, error) {

	result := &response.ResponseCalendarAdd{}

	options := &object.HashMap{
		"calendar": calendar,
		"agentid":  agentID,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/calendar/add", options, nil, nil, result)

	return result, err
}

// 更新日历
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func (comp *Client) Update(calendar *power.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"calendar": calendar,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/calendar/update", options, nil, nil, result)

	return result, err
}

// 获取日历详情
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func (comp *Client) Get(calIDList []string) (*response.ResponseCalendarGet, error) {

	result := &response.ResponseCalendarGet{}

	options := &object.HashMap{
		"cal_id_list": calIDList,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/calendar/get", options, nil, nil, result)

	return result, err
}

// 删除日历
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func (comp *Client) Del(calID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"cal_id": calID,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/calendar/del", options, nil, nil, result)

	return result, err
}
