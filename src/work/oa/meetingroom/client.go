package meetingroom

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/meetingroom/request"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/meetingroom/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 添加会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func (comp *Client) Add(options *request.RequestMeetingRoomAdd) (*response.ResponseMeetingRoomAdd, error) {

	result := &response.ResponseMeetingRoomAdd{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/add", options, nil, nil, result)

	return result, err
}


// 查询会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func (comp *Client) List(options *request.RequestMeetingRoomList) (*response.ResponseMeetingRoomList, error) {

	result := &response.ResponseMeetingRoomList{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/list", options, nil, nil, result)

	return result, err
}

// 编辑会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func (comp *Client) Edit(options *request.RequestMeetingRoomEdit) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/edit", options, nil, nil, result)

	return result, err
}

// 删除会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func (comp *Client) Del(options *request.RequestMeetingRoomDel) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/del", options, nil, nil, result)

	return result, err
}


// 查询会议室的预定信息
// https://work.weixin.qq.com/api/doc/90000/90135/93620
func (comp *Client) GetBookingInfo(options *request.RequestMeetingRoomGetBookingInfo) (*response.ResponseMeetingRoomGetBookingInfo, error) {

	result := &response.ResponseMeetingRoomGetBookingInfo{}

	_, err := comp.HttpPostJson("cgi-bin/oa/meetingroom/get_booking_info", options, nil, nil, result)

	return result, err
}