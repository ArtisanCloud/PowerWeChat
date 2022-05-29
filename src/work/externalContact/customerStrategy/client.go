package customerStrategy

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/externalContact/customerStrategy/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/externalContact/customerStrategy/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 获取规则组列表
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func (comp *Client) List(cursor string, limit int64) (*response.ResponseCustomerStrategyList, error) {

	result := &response.ResponseCustomerStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/list", options, nil, nil, result)

	return result, err
}

// 获取规则组详情
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func (comp *Client) Get(strategyID int64) (*response.ResponseCustomerStrategyGet, error) {

	result := &response.ResponseCustomerStrategyGet{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/get", options, nil, nil, result)

	return result, err
}

// 获取规则组管理范围
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func (comp *Client) GetRange(strategyID int64, cursor string, limit int64) (*response.ResponseCustomerStrategyGetRange, error) {

	result := &response.ResponseCustomerStrategyGetRange{}

	options := &object.HashMap{
		"strategy_id": strategyID,
		"cursor":      cursor,
		"limit":       limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/get_range", options, nil, nil, result)

	return result, err
}

// 创建新的规则组
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func (comp *Client) Create(options *request.RequestCustomerStrategyCreate) (*response.ResponseCustomerStrategyCreate, error) {

	result := &response.ResponseCustomerStrategyCreate{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/create", options, nil, nil, result)

	return result, err
}

// 编辑规则组及其管理范围
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func (comp *Client) Edit(options *request.RequestCustomerStrategyEdit) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/edit", options, nil, nil, result)

	return result, err
}

// 删除规则组
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func (comp *Client) Del(strategyID int64) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/del", options, nil, nil, result)

	return result, err
}
