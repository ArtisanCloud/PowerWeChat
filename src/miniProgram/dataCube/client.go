package dataCube

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/miniprogram/dataCube/response"
)

type Client struct {
	*kernel.BaseClient
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
func (comp *Client) SummaryTrend(from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query("datacube/getweanalysisappiddailysummarytrend", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html
func (comp *Client) DailyVisitTrend(from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query("datacube/getweanalysisappiddailyvisittrend", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html
func (comp *Client) WeeklyVisitTrend(from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query("datacube/getweanalysisappidweeklyvisittrend", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html
func (comp *Client) MonthlyVisitTrend(from string, to string) (*response.ResponseDataCubeSummary, error) {

	result := &response.ResponseDataCubeSummary{}

	_, err := comp.Query("datacube/getweanalysisappidmonthlyvisittrend", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html
func (comp *Client) VisitDistribution(from string, to string) (*response.ResponseDataCubeVisit, error) {

	result := &response.ResponseDataCubeVisit{}

	_, err := comp.Query("datacube/getweanalysisappidvisitdistribution", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func (comp *Client) DailyRetainInfo(from string, to string) (*response.ResponseDataCubeVisitInfo, error) {

	result := &response.ResponseDataCubeVisitInfo{}

	_, err := comp.Query("datacube/getweanalysisappiddailyretaininfo", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
func (comp *Client) WeeklyRetainInfo(from string, to string) (*response.ResponseDataCubeVisitInfo, error) {

	result := &response.ResponseDataCubeVisitInfo{}

	_, err := comp.Query("datacube/getweanalysisappidweeklyretaininfo", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
func (comp *Client) MonthlyRetainInfo(from string, to string) (*response.ResponseDataCubeVisitInfo, error) {

	result := &response.ResponseDataCubeVisitInfo{}

	_, err := comp.Query("datacube/getweanalysisappidmonthlyretaininfo", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html
func (comp *Client) VisitPage(from string, to string) (*response.ResponseDataCubeVisit, error) {

	result := &response.ResponseDataCubeVisit{}

	_, err := comp.Query("datacube/getweanalysisappidvisitpage", from, to, nil, result)

	return result, err
}

// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func (comp *Client) UserPortrait(from string, to string) (*response.ResponseDataCubeVisitInfo, error) {

	result := &response.ResponseDataCubeVisitInfo{}

	_, err := comp.Query("datacube/getweanalysisappiduserportrait", from, to, nil, result)

	return result, err
}

func (comp *Client) Query(endpoint string, from string, to string, outHeader interface{}, outBody interface{}) (interface{}, error) {

	data := &object.HashMap{
		"begin_date": from,
		"end_date":   to,
	}

	rs, err := comp.HttpPostJson(endpoint, data, nil, outHeader, outBody)

	return rs, err
}
