package statistics

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/statistics/request"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/statistics/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取「联系客户统计」数据
// https://work.weixin.qq.com/api/doc/90000/90135/92132
func (comp *Client) GetUserBehaviorData(options *request.RequestGetUserBehaviorData) (*response.ResponseGetUserBehaviorData, error) {

	result := &response.ResponseGetUserBehaviorData{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/get_user_behavior_data", options, nil, nil, result)

	return result, err
}



// 获取「群聊数据统计」数据
// https://work.weixin.qq.com/api/doc/90000/90135/92133
func (comp *Client) Statistic(options *request.RequestStatistic) (*response.ResponseStatistic, error) {

	result := &response.ResponseStatistic{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/groupchat/statistic", options, nil, nil, result)

	return result, err
}
