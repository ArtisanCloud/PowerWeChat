package customerStrategy

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/customerStrategy/response"
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
func (comp *Client) List(cursor string, limit int) (*response.ResponseCustomerStrategyList, error) {

	result := &response.ResponseCustomerStrategyList{}

	options := &object.HashMap{
		"cursor": cursor,
		"limit":  limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/externalcontact/customer_strategy/list", options, nil, nil, result)

	return result, err
}
