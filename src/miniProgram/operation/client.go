package operation

import (
	response2 "github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/operation/response"
)

type Client struct {
	*kernel.BaseClient
}

// 查询域名配置
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getDomainInfo.html
func (comp *Client) GetDomainInfo(action string) (*response.ResponseOperationGetDomainInfo, error) {

	result := &response.ResponseOperationGetDomainInfo{}

	data := &object.HashMap{
		"action": action,
	}

	_, err := comp.HttpPostJson("wxa/getwxadevinfo", data, nil, nil, result)

	return result, err
}

// 获取用户反馈列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedback.html
func (comp *Client) GetFeedback(feedbackType string, page int, num int) (*response.ResponseOperationGetFeedback, error) {

	result := &response.ResponseOperationGetFeedback{}

	params := &object.HashMap{
		"type": feedbackType,
		"page": page,
		"num":  num,
	}

	_, err := comp.HttpGet("wxaapi/feedback/list", params, nil, result)

	return result, err
}

// 获取 mediaId 图片
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedbackmedia.html
func (comp *Client) GetFeedbackMedia(recordID string, mediaID string) (*response2.HttpResponse, error) {

	params := &object.HashMap{
		"record_id": recordID,
		"media_id":  mediaID,
	}

	rs, err := comp.HttpGet("cgi-bin/media/getfeedbackmedia", params, nil, nil)

	return rs.(*response2.HttpResponse), err
}


// 查询当前分阶段发布详情
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getGrayReleasePlan.html
func (comp *Client) GetGrayReleasePlan() (*response.ResponseOperationGetGrayReleasePlan, error) {

	result := &response.ResponseOperationGetGrayReleasePlan{}


	_, err := comp.HttpGet("wxa/getgrayreleaseplan", nil, nil, result)

	return result, err
}


// 错误查询详情
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrDetail.html
func (comp *Client) GetJsErrDetail(data *power.HashMap) (*response.ResponseOperationGetJsErrDetail, error) {

	result := &response.ResponseOperationGetJsErrDetail{}

	_, err := comp.HttpPostJson("wxaapi/log/jserr_detail", data, nil,nil, result)

	return result, err
}

// 错误查询列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrList.html
func (comp *Client) GetJsErrList(data *power.HashMap) (*response.ResponseOperationGetJsErrList, error) {

	result := &response.ResponseOperationGetJsErrList{}

	_, err := comp.HttpPostJson("wxaapi/log/jserr_list", data, nil,nil, result)

	return result, err
}


// 错误查询, 接口即将废弃
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrSearch.html
func (comp *Client) GetJsErrSearch(data *power.HashMap) (*response.ResponseOperationGetJsErrSearch, error) {

	result := &response.ResponseOperationGetJsErrSearch{}

	_, err := comp.HttpPostJson("wxaapi/log/jserr_search", data, nil,nil, result)

	return result, err
}


// 性能监控
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getPerformance.html
func (comp *Client) GetPerformance(data *power.HashMap) (*response.ResponseOperationGetPerformance, error) {

	result := &response.ResponseOperationGetPerformance{}

	_, err := comp.HttpPostJson("wxaapi/log/get_performance", data, nil,nil, result)

	return result, err
}


// 获取访问来源
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getSceneList.html
func (comp *Client) GetSceneList() (*response.ResponseOperationGetSceneList, error) {

	result := &response.ResponseOperationGetSceneList{}


	_, err := comp.HttpGet("wxa/log/get_scene", nil, nil, result)

	return result, err
}

// 获取客户端版本
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getVersionList.html
func (comp *Client) GetVersionList() (*response.ResponseOperationGetVersionList, error) {

	result := &response.ResponseOperationGetVersionList{}

	_, err := comp.HttpGet("wxaapi/log/get_client_version", nil, nil, result)

	return result, err
}

// 实时日志查询
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.realtimelogSearch.html
func (comp *Client) RealTimeLogSearch(params *object.StringMap) (*response.ResponseOperationRealTimeLogSearch, error) {

	result := &response.ResponseOperationRealTimeLogSearch{}

	_, err := comp.HttpGet("wxaapi/userlog/userlog_search", params, nil, result)

	return result, err
}