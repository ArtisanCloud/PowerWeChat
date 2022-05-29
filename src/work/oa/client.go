package oa

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/oa/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/oa/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取企业所有打卡规则
// https://work.weixin.qq.com/api/doc/90000/90135/93384
func (comp *Client) GetCorpCheckInOption() (*response.ResponseCorpCheckInRules, error) {
	result := &response.ResponseCorpCheckInRules{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcorpcheckinoption", nil, nil, nil, result)

	return result, err

}

// 获取员工打卡规则
// https://work.weixin.qq.com/api/doc/90000/90135/90263
func (comp *Client) GetCheckInOption(datetime int, userList []string) (*response.ResponseCheckInRules, error) {
	result := &response.ResponseCheckInRules{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckinoption", &object.HashMap{
		"datetime":   fmt.Sprintf("%d", datetime),
		"userIDlist": userList,
	}, nil, nil, result)

	return result, err
}

// 获取打卡记录数据
// https://work.weixin.qq.com/api/doc/90000/90135/90262
func (comp *Client) GetCheckinData(recordType int, startTime int, endTime int, userList []string) (*response.ResponseCheckInRecords, error) {
	result := &response.ResponseCheckInRecords{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckindata", &object.HashMap{
		"opencheckindatatype": fmt.Sprintf("%d", recordType),
		"starttime":           fmt.Sprintf("%d", startTime),
		"endtime":             fmt.Sprintf("%d", endTime),
		"userIDlist":          userList,
	}, nil, nil, result)

	return result, err
}

// 获取打卡日报数据
// https://work.weixin.qq.com/api/doc/90000/90135/93374
func (comp *Client) GetCheckinDayData(startTime int, endTime int, userIDs []string) (*response.ResponseCheckInDatas, error) {

	result := &response.ResponseCheckInDatas{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckin_daydata", &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"userIDlist": userIDs,
	}, nil, nil, result)

	return result, err

}

// 获取打卡月报数据
// https://work.weixin.qq.com/api/doc/90000/90135/93387
func (comp *Client) GetCheckInMonthData(startTime int, endTime int, userIDs []string) (*response.ResponseCheckInDatas, error) {
	result := &response.ResponseCheckInDatas{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckin_monthdata", &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"userIDlist": userIDs,
	}, nil, nil, result)

	return result, err
}

// 获取打卡人员排班信息
// https://work.weixin.qq.com/api/doc/90000/90135/93380
func (comp *Client) GetCheckInScheduleList(startTime int, endTime int, userIDs []string) (*response.ResponseCheckInSchedulist, error) {
	result := &response.ResponseCheckInSchedulist{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/getcheckinschedulist", &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"userIDlist": userIDs,
	}, nil, nil, result)

	return result, err
}

// 为打卡人员排班
// https://work.weixin.qq.com/api/doc/90000/90135/93385
func (comp *Client) SetCheckInScheduleList(options *request.RequestCheckInSetScheduleList) (*response2.ResponseWork, error) {
	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/setcheckinschedulist", options, nil, nil, result)

	return result, err
}

// 录入打卡人员人脸信息
// https://work.weixin.qq.com/api/doc/90000/90135/93378
func (comp *Client) AddCheckInUserFace(userID string, userFace string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/checkin/addcheckinuserface", &object.HashMap{
		"userID":   userID,
		"userface": userFace,
	}, nil, nil, result)

	return result, err
}

// 获取审批模板详情
// https://work.weixin.qq.com/api/doc/90000/90135/91982
func (comp *Client) ApprovalTemplate(templateID string) (*response.ResponseApprovalTemplate, error) {
	result := &response.ResponseApprovalTemplate{}

	_, err := comp.HttpPostJson("cgi-bin/oa/gettemplatedetail", &object.HashMap{
		"template_id": templateID,
	}, nil, nil, result)

	return result, err
}

// 提交审批申请
// https://work.weixin.qq.com/api/doc/90000/90135/91853
func (comp *Client) CreateApproval(data *power.HashMap) (*response.ResponseApprovalCreate, error) {
	result := &response.ResponseApprovalCreate{}

	_, err := comp.HttpPostJson("cgi-bin/oa/applyevent", data, nil, nil, result)

	return result, err
}

// 批量获取审批单号
// https://work.weixin.qq.com/api/doc/90000/90135/91816
func (comp *Client) GetApprovalInfo(startTime int, endTime int, nextCursor int, size int, filters *object.HashMap) (*response.ResponseApprovalNoList, error) {

	result := &response.ResponseApprovalNoList{}

	if size > 100 {
		size = 100
	}

	options := &object.HashMap{
		"starttime": fmt.Sprintf("%d", startTime),
		"endtime":   fmt.Sprintf("%d", endTime),
		"cursor":    fmt.Sprintf("%d", nextCursor),
		"size":      size,
		"filters":   filters,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/getapprovalinfo", options, nil, nil, result)

	return result, err
}

// 获取审批申请详情
// https://work.weixin.qq.com/api/doc/90000/90135/91983
func (comp *Client) GetApprovalDetail(number int) (*response.ResponseApprovalDetail, error) {
	result := &response.ResponseApprovalDetail{}

	_, err := comp.HttpPostJson("cgi-bin/oa/getapprovaldetail", &object.HashMap{
		"sp_no": fmt.Sprintf("%d", number),
	}, nil, nil, result)

	return result, err
}

// 获取审批数据（旧）
// https://work.weixin.qq.com/api/doc/90000/90135/91530
func (comp *Client) GetApprovalData(startTime int, endTime int, nextNumber int) (*response.ResponseApprovalGetData, error) {

	result := &response.ResponseApprovalGetData{}

	options := &object.HashMap{
		"starttime":  fmt.Sprintf("%d", startTime),
		"endtime":    fmt.Sprintf("%d", endTime),
		"next_spnum": fmt.Sprintf("%d", nextNumber),
	}
	_, err := comp.HttpPostJson("cgi-bin/corp/getapprovaldata", options, nil, nil, result)

	return result, err

}

// 获取企业假期管理配置
// https://work.weixin.qq.com/api/doc/90000/90135/93375
func (comp *Client) GetCorpConfig() (*response.ResponseCorpVacationGetConfig, error) {

	result := &response.ResponseCorpVacationGetConfig{}

	_, err := comp.HttpGet("cgi-bin/corp/vacation/getcorpconf", nil, nil, result)

	return result, err

}

// 获取成员假期余额
// https://work.weixin.qq.com/api/doc/90000/90135/93376
func (comp *Client) GetUserVacationQuota(userID string) (*response.ResponseCorpVacationGetQuota, error) {

	result := &response.ResponseCorpVacationGetQuota{}

	_, err := comp.HttpPostJson("cgi-bin/corp/vacation/getuservacationquota", &object.HashMap{
		"userid": userID,
	}, nil, nil, result)

	return result, err

}

// 修改成员假期余额
// https://work.weixin.qq.com/api/doc/90000/90135/93376
func (comp *Client) SetOneUserQuota(data *object.HashMap) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/corp/vacation/setoneuserquota", data, nil, nil, result)

	return result, err

}
