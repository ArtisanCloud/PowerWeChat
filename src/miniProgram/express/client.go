package express

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/express/response"
)

type Client struct {
	*kernel.BaseClient
}

// 生成运单
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.addOrder.html
func (comp *Client) AddOrder(data *object.HashMap) (*response.ResponseExpressAddOrder, error) {

	result := &response.ResponseExpressAddOrder{}

	_, err := comp.HttpPostJson("cgi-bin/express/business/order/add", data, nil, nil, result)

	return result, err
}


// 批量获取运单数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.batchGetOrder.html
func (comp *Client) BatchGetOrder(orderList []*object.HashMap) (*response.ResponseExpressBatchOrderList, error) {

	result := &response.ResponseExpressBatchOrderList{}

	data:=&object.HashMap{
		"order_list": orderList,
	}

	_, err := comp.HttpPostJson("cgi-bin/express/business/order/batchget", data, nil, nil, result)

	return result, err
}
