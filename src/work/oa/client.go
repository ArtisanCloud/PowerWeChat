package oa

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90262

func (comp *Client) CheckInRecords(startTime int, endTime int, userList []string, recordType int) (*response.ResponseCheckInRecords, error) {
	result := &response.ResponseCheckInRecords{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckindata", &object.HashMap{
		"opencheckindatatype": fmt.Sprintf("%d", recordType),
		"starttime":           fmt.Sprintf("%d", startTime),
		"endtime":             fmt.Sprintf("%d", endTime),
		"userIDlist":          userList,
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/90263
func (comp *Client) CheckInRules(datetime int, userList []string) (*response.ResponseCheckInRules, error) {
	result := &response.ResponseCheckInRules{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckinoption", &object.HashMap{
		"datetime":   fmt.Sprintf("%d", datetime),
		"userIDlist": userList,
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93384
func (comp *Client) CorpCheckInRules() (*response.ResponseCorpCheckInRules, error) {
	result := &response.ResponseCorpCheckInRules{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcorpcheckinoption", nil, nil, nil, result)

	return result, err

}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93374
func (comp *Client) CheckInDayData(startTime int, endTime int, userIDs []string) (*response.ResponseCheckInDatas, error) {

	result := &response.ResponseCheckInDatas{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckin_daydata", &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"userIDlist": userIDs,
	}, nil, nil, result)

	return result, err

}
func (comp *Client) CheckInMonthData(startTime int, endTime int, userIDs []string) (*response.ResponseCheckInDatas, error) {
	result := &response.ResponseCheckInDatas{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckin_monthdata", &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"userIDlist": userIDs,
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93380
func (comp *Client) CheckInSchedules(startTime int, endTime int, userIDs []string) (*response.ResponseCheckInSchedulist, error) {
	result := &response.ResponseCheckInSchedulist{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckin_monthdata", &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"userIDlist": userIDs,
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93385
func (comp *Client) SetCheckInSchedules(params *power.HashMap) (*response2.ResponseWork, error) {
	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/setcheckinschedulist", params, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93378
func (comp *Client) AddCheckInUserFace(userID string, userFace string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/addcheckinuserface", &object.HashMap{
		"userID":   userID,
		"userface": userFace,
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91982
func (comp *Client) ApprovalTemplate(templateID string) (*response.ResponseApprovalTemplate, error) {
	result := &response.ResponseApprovalTemplate{}

	_, err := comp.HttpPostJson("cgi-bin/oa/gettemplatedetail", &object.HashMap{
		"template_id": templateID,
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91853
func (comp *Client) CreateApproval(data *power.HashMap) (*response.ResponseApprovalCreate, error) {
	result := &response.ResponseApprovalCreate{}

	_, err := comp.HttpPostJson("cgi-bin/oa/applyevent", data, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91816
func (comp *Client) GetApprovalInfo(startTime int, endTime int, nextCursor int, size int, filters *object.HashMap) (*response.ResponseApprovalNoList, error) {

	result := &response.ResponseApprovalNoList{}

	if size > 100 {
		size = 100
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/getapprovalinfo", &object.HashMap{
		"starttime": fmt.Sprintf("%d", startTime),
		"endtime":   fmt.Sprintf("%d", endTime),
		"cursor":    fmt.Sprintf("%d", nextCursor),
		"size":      size,
		"filters":   filters,
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91983
func (comp *Client) ApprovalDetail(number int) (*response.ResponseApprovalDetail, error) {
	result := &response.ResponseApprovalDetail{}

	_, err := comp.HttpPostJson("cgi-bin/oa/getapprovaldetail", &object.HashMap{
		"sp_no": fmt.Sprintf("%d", number),
	}, nil, nil, result)

	return result, err
}

// https://open.work.weixin.qq.com/api/doc/90000/90135/91530
func (comp *Client) ApprovalRecords(startTime int, endTime int, nextNumber int) (*response.ResponseApprovalGetData, error) {

	result := &response.ResponseApprovalGetData{}

	_, err := comp.HttpPostJson("cgi-bin/corp/getapprovaldata", &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"next_spnum": fmt.Sprintf("%d", nextNumber),
	}, nil, nil, result)

	return result, err

}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93375
func (comp *Client) GetCorpConfig() (*response.ResponseCorpVacationGetConfig, error) {

	result := &response.ResponseCorpVacationGetConfig{}

	_, err := comp.HttpGet("cgi-bin/corp/vacation/getcorpconf", nil, nil, result)

	return result, err

}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93376
func (comp *Client) GetUserVacationQuota(userID string) (*response.ResponseCorpVacationGetQuota, error) {

	result := &response.ResponseCorpVacationGetQuota{}

	_, err := comp.HttpPostJson("cgi-bin/corp/vacation/getuservacationquota", &object.HashMap{
		"userid": userID,
	}, nil, nil, result)

	return result, err

}

// https://open.work.weixin.qq.com/api/doc/90000/90135/93376
func (comp *Client) SeToneUserQuota(data *object.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/corp/vacation/setoneuserquota", data, nil, nil, result)

	return result, err

}
