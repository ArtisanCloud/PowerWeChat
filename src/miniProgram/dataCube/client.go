package dataCube

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/dataCube/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/dataCube/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取用户访问小程序数据概况
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
func (comp *Client) GetDailySummary(ctx context.Context, from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappiddailysummarytrend", from, to, nil, result)

	return result, err
}

// 获取小程序启动性能，运行性能等数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getPerformanceData.html
func (comp *Client) GetPerformanceData(ctx context.Context, options *request.RequestGetPerformanceData) (*response.ResponseDataCubeGetPerformanceData, error) {

	result := &response.ResponseDataCubeGetPerformanceData{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/performance/boot", options, nil, nil, result)

	return result, err
}

// 获取用户访问小程序数据日趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html
func (comp *Client) GetDailyVisitTrend(ctx context.Context, from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappiddailyvisittrend", from, to, nil, result)

	return result, err
}

// 获取用户访问小程序数据周趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html
func (comp *Client) GetWeeklyVisitTrend(ctx context.Context, from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappidweeklyvisittrend", from, to, nil, result)

	return result, err
}

// 获取用户访问小程序数据月趋势(能查询到的最新数据为上一个自然月的数据)
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html
func (comp *Client) GetMonthlyVisitTrend(ctx context.Context, from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappidmonthlyvisittrend", from, to, nil, result)

	return result, err
}

// 获取用户小程序访问分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html
func (comp *Client) GetVisitDistribution(ctx context.Context, from string, to string) (*response.ResponseDataCubeVisit, error) {

	result := &response.ResponseDataCubeVisit{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappidvisitdistribution", from, to, nil, result)

	return result, err
}

// 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func (comp *Client) GetDailyRetainInfo(ctx context.Context, from string, to string) (*response.ResponseDataCubeVisitInfo, error) {

	result := &response.ResponseDataCubeVisitInfo{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappiddailyretaininfo", from, to, nil, result)

	return result, err
}

// 获取用户访问小程序周留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
func (comp *Client) GetWeeklyRetainInfo(ctx context.Context, from string, to string) (*response.ResponseDataCubeVisitInfo, error) {

	result := &response.ResponseDataCubeVisitInfo{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappidweeklyretaininfo", from, to, nil, result)

	return result, err
}

// 获取用户访问小程序月留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
func (comp *Client) GetMonthlyRetainInfo(ctx context.Context, from string, to string) (*response.ResponseDataCubeVisitInfo, error) {

	result := &response.ResponseDataCubeVisitInfo{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappidmonthlyretaininfo", from, to, nil, result)

	return result, err
}

// 访问页面
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html
func (comp *Client) GetVisitPage(ctx context.Context, from string, to string) (*response.ResponseDataCubeVisit, error) {

	result := &response.ResponseDataCubeVisit{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappidvisitpage", from, to, nil, result)

	return result, err
}

// 获取小程序新增或活跃用户的画像分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func (comp *Client) GetUserPortrait(ctx context.Context, from string, to string) (*response.ResponseDataCubeUserPortrait, error) {

	result := &response.ResponseDataCubeUserPortrait{}

	_, err := comp.Query(ctx, "datacube/getweanalysisappiduserportrait", from, to, nil, result)

	return result, err
}

func (comp *Client) Query(ctx context.Context, endpoint string, from string, to string, outHeader interface{}, outBody interface{}) (interface{}, error) {

	data := &object.HashMap{
		"begin_date": from,
		"end_date":   to,
	}

	rs, err := comp.BaseClient.HttpPostJson(ctx, endpoint, data, nil, outHeader, outBody)

	return rs, err
}
