package store

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/store/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 拉取门店小程序类目
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Categories() (*response.ResponseStoreCategory, error) {

	result := &response.ResponseStoreCategory{}

	_, err := comp.HttpGet("wxa/get_merchant_category", nil, nil, result)

	return result, err
}
