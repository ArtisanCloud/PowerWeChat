package momentStrategy

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/momentStrategy/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/momentStrategy/response"
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

// 获取规则组列表
// https://developer.work.weixin.qq.com/document/path/94890
func (comp *Client) List(ctx context.Context, cursor string, limit int) (*response.ResponseMomentStrategyList, error) {

	result := &response.ResponseMomentStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/moment_strategy/list", options, nil, nil, result)

	return result, err
}

// 获取规则组详情
// https://developer.work.weixin.qq.com/document/path/94890
func (comp *Client) Get(ctx context.Context, strategyID int) (*response.ResponseMomentStrategyGet, error) {

	result := &response.ResponseMomentStrategyGet{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/moment_strategy/list", options, nil, nil, result)

	return result, err
}

// 获取规则组管理范围
// https://developer.work.weixin.qq.com/document/path/94890
func (comp *Client) GetRange(ctx context.Context, strategyID int, cursor string, limit int) (*response.ResponseMomentStrategyGetRange, error) {

	result := &response.ResponseMomentStrategyGetRange{}

	options := &object.HashMap{
		"strategy_id": strategyID,
		"cursor":      cursor,
		"limit":       limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/moment_strategy/get_range", options, nil, nil, result)

	return result, err
}

// 创建新的规则组
// https://developer.work.weixin.qq.com/document/path/94890
func (comp *Client) Create(ctx context.Context, options *request.RequestMomentStrategyCreate) (*response.ResponseMomentStrategyCreate, error) {

	result := &response.ResponseMomentStrategyCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/moment_strategy/create", options, nil, nil, result)

	return result, err
}

// 编辑规则组及其管理范围
// https://developer.work.weixin.qq.com/document/path/94890
func (comp *Client) Edit(ctx context.Context, options *request.RequestMomentStrategyEdit) (*response.ResponseMomentStrategyCreate, error) {

	result := &response.ResponseMomentStrategyCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/moment_strategy/edit", options, nil, nil, result)

	return result, err
}

// 删除规则组
// https://developer.work.weixin.qq.com/document/path/94890
func (comp *Client) Del(ctx context.Context, strategyID int) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/moment_strategy/del", options, nil, nil, result)

	return result, err
}
