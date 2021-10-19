package momentStrategy

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel"
	response2 "github.com/ArtisanCloud/powerwechat/src/kernel/response"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/momentStrategy/request"
	"github.com/ArtisanCloud/powerwechat/src/work/externalContact/momentStrategy/response"
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
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *Client) List(cursor string, limit int) (*response.ResponseMomentStrategyList, error) {

	result := &response.ResponseMomentStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/list", options, nil, nil, result)

	return result, err
}

// 获取规则组详情
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *Client) Get(strategyID int) (*response.ResponseMomentStrategyGet, error) {

	result := &response.ResponseMomentStrategyGet{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/list", options, nil, nil, result)

	return result, err
}

// 获取规则组管理范围
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *Client) GetRange(strategyID int, cursor string, limit int) (*response.ResponseMomentStrategyGetRange, error) {

	result := &response.ResponseMomentStrategyGetRange{}

	options := &object.HashMap{
		"strategy_id": strategyID,
		"cursor":      cursor,
		"limit":       limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/get_range", options, nil, nil, result)

	return result, err
}


// 创建新的规则组
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *Client) Create(options *request.RequestMomentStrategyCreate) (*response.ResponseMomentStrategyCreate, error) {

	result := &response.ResponseMomentStrategyCreate{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/create", options, nil, nil, result)

	return result, err
}


// 编辑规则组及其管理范围
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *Client) Edit(options *request.RequestMomentStrategyEdit) (*response.ResponseMomentStrategyCreate, error) {

	result := &response.ResponseMomentStrategyCreate{}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/edit", options, nil, nil, result)

	return result, err
}


// 删除规则组
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func (comp *Client) Del(strategyID int) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	options := &object.HashMap{
		"strategy_id": strategyID,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/moment_strategy/del", options, nil, nil, result)

	return result, err
}