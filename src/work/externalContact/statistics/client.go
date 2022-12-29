package statistics

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/statistics/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/statistics/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 获取「联系客户统计」数据
// https://developer.work.weixin.qq.com/document/path/92132
func (comp *Client) GetUserBehaviorData(ctx context.Context, options *request.RequestGetUserBehaviorData) (*response.ResponseGetUserBehaviorData, error) {

	result := &response.ResponseGetUserBehaviorData{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/get_user_behavior_data", options, nil, nil, result)

	return result, err
}

// 获取「群聊数据统计」数据
// https://developer.work.weixin.qq.com/document/path/92133
func (comp *Client) Statistic(ctx context.Context, options *request.RequestStatistic) (*response.ResponseStatistic, error) {

	result := &response.ResponseStatistic{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/groupchat/statistic", options, nil, nil, result)

	return result, err
}
