package meeting

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/oa/meeting/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/oa/meeting/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 创建预约会议
// https://work.weixin.qq.com/api/doc/90000/90135/93627
func (comp *Client) Create(options *request.RequestMeetingCreate) (*response.ResponseMeetingCreate, error) {

	result := &response.ResponseMeetingCreate{}

	_, err := comp.HttpPostJson("cgi-bin/meeting/create", options, nil, nil, result)

	return result, err
}

// 修改预约会议
// https://work.weixin.qq.com/api/doc/90000/90135/93631
func (comp *Client) Update(options *request.RequestMeetingUpdate) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/meeting/update", options, nil, nil, result)

	return result, err
}

// 取消预约会议
// https://work.weixin.qq.com/api/doc/90000/90135/93630
func (comp *Client) Cancel(meetingID string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"meetingid": meetingID,
	}

	_, err := comp.HttpPostJson("cgi-bin/meeting/cancel", options, nil, nil, result)

	return result, err
}

// 获取成员会议ID列表
// https://work.weixin.qq.com/api/doc/90000/90135/93628
func (comp *Client) GetUserMeetingID(userID string, cursor string, beginTime int64, endTime int64, limit int) (*response.ResponseMeetingGetUserMeetingID, error) {

	result := &response.ResponseMeetingGetUserMeetingID{}

	options := &object.HashMap{
		"userid":     userID,
		"cursor":     cursor,
		"begin_time": beginTime,
		"end_time":   endTime,
		"limit":      limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/meeting/get_user_meetingid", options, nil, nil, result)

	return result, err
}

// 获取会议详情
// https://work.weixin.qq.com/api/doc/90000/90135/93629
func (comp *Client) GetBookingInfo(meetingID string) (*response.ResponseMeetingGetBookingInfo, error) {

	result := &response.ResponseMeetingGetBookingInfo{}

	options := &object.HashMap{
		"meetingid": meetingID,
	}

	_, err := comp.HttpPostJson("cgi-bin/meeting/get_info", options, nil, nil, result)

	return result, err
}
