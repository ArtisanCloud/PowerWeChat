package customerStrategy

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/customerStrategy/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/customerStrategy/response"
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
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) List(ctx context.Context, cursor string, limit int64) (*response.ResponseCustomerStrategyList, error) {

	result := &response.ResponseCustomerStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_strategy/list", options, nil, nil, result)

	return result, err
}

// 获取规则组详情
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Get(ctx context.Context, strategyID int64) (*response.ResponseCustomerStrategyGet, error) {

	result := &response.ResponseCustomerStrategyGet{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_strategy/get", options, nil, nil, result)

	return result, err
}

// 获取规则组管理范围
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) GetRange(ctx context.Context, strategyID int64, cursor string, limit int64) (*response.ResponseCustomerStrategyGetRange, error) {

	result := &response.ResponseCustomerStrategyGetRange{}

	options := &object.HashMap{
		"strategy_id": strategyID,
		"cursor":      cursor,
		"limit":       limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_strategy/get_range", options, nil, nil, result)

	return result, err
}

// 创建新的规则组
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Create(ctx context.Context, options *request.RequestCustomerStrategyCreate) (*response.ResponseCustomerStrategyCreate, error) {

	result := &response.ResponseCustomerStrategyCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_strategy/create", options, nil, nil, result)

	return result, err
}

// 编辑规则组及其管理范围
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Edit(ctx context.Context, options *request.RequestCustomerStrategyEdit) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_strategy/edit", options, nil, nil, result)

	return result, err
}

// 删除规则组
// https://developer.work.weixin.qq.com/document/path/94883
func (comp *Client) Del(ctx context.Context, strategyID int64) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/externalcontact/customer_strategy/del", options, nil, nil, result)

	return result, err
}
