package dataCube

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
	response2 "github.com/ArtisanCloud/power-wechat/src/miniProgram/express/response"
)

type Client struct {
	*kernel.BaseClient
}

// 异常件退回商家商家确认收货接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.abnormalConfirm.html
func (comp *Client) AbnormalConfirm(options *object.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/order/confirm_return", options, nil, nil, result)

	return result, err
}

// 订单增加小费
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addOrder.html
func (comp *Client) AddOrder(options *object.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/order/add", options, nil, nil, result)

	return result, err
}

// 订单增加小费
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addTip.html
func (comp *Client) AddTips(options *object.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/order/addtips", options, nil, nil, result)

	return result, err
}

// 第三方代商户发起绑定配送公司帐号的请求
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.bindAccount.html
func (comp *Client) AddShops(deliveryID string) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	query := object.StringMap{
		"delivery_id": deliveryID,
	}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/order/addtips", nil, query, nil, result)

	return result, err
}

// 取消配送单接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.cancelOrder.html
func (comp *Client) CancelOrder(options *object.HashMap) (*response2.ResponseExpressCancelOrder, error) {

	result := &response2.ResponseExpressCancelOrder{}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/order/cancel", options, nil, nil, result)

	return result, err
}

// 获取已支持的配送公司列表接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getAllImmeDelivery.html
func (comp *Client) GetAllImmeDelivery() (*response2.ResponseExpressDelivery, error) {

	result := &response2.ResponseExpressDelivery{}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/delivery/getall", nil, nil, nil, result)

	return result, err
}

// 拉取已绑定账号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getAllImmeDelivery.html
func (comp *Client) GetBindAccount() (*response2.ResponseExpressBindAccount, error) {

	result := &response2.ResponseExpressBindAccount{}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/delivery/getall", nil, nil, nil, result)

	return result, err
}

// 拉取配送单信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getOrder.html
func (comp *Client) GetOrder(shopID string, shopOrderID string, shopNO string, deliverySign string) (*response2.ResponseExpressGetOrder, error) {

	result := &response2.ResponseExpressGetOrder{}

	data := object.StringMap{
		"shopid":        shopID,
		"shop_order_id": shopOrderID,
		"shop_no":       shopNO,
		"delivery_sign": deliverySign,
	}

	_, err := comp.HttpPostJson("cgi-bin/express/local/business/order/addtips", data, nil, nil, result)

	return result, err
}
