package immediateDelivery

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/immediateDelivery/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 异常件退回商家商家确认收货接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.abnormalConfirm.html
func (comp *Client) AbnormalConfirm(ctx context.Context, options *power.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/local/business/order/confirm_return", options, nil, nil, result)

	return result, err
}

// 订单增加小费
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addOrder.html
func (comp *Client) AddOrder(ctx context.Context, options *power.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/add", options, nil, nil, result)

	return result, err
}

// 订单增加小费
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addTip.html
func (comp *Client) AddTips(ctx context.Context, options *power.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/addtips", options, nil, nil, result)

	return result, err
}

// 第三方代商户发起绑定配送公司帐号的请求
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.bindAccount.html
func (comp *Client) BindAccount(ctx context.Context, deliveryID string) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	options := &object.HashMap{
		"delivery_id": deliveryID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/local/business/shop/add", options, nil, nil, result)

	return result, err
}

// 第三方代商户发起绑定配送公司帐号的请求
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.bindAccount.html
func (comp *Client) AddShops(ctx context.Context, deliveryID string) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	query := &object.StringMap{
		"delivery_id": deliveryID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/addtips", nil, query, nil, result)

	return result, err
}

// 取消配送单接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.cancelOrder.html
func (comp *Client) CancelOrder(ctx context.Context, options *power.HashMap) (*response2.ResponseImmediateDeliveryCancelOrder, error) {

	result := &response2.ResponseImmediateDeliveryCancelOrder{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/cancel", options, nil, nil, result)

	return result, err
}

// 获取已支持的配送公司列表接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getAllImmeDelivery.html
func (comp *Client) GetAllImmeDelivery(ctx context.Context) (*response2.ResponseImmediateDeliveryDelivery, error) {

	result := &response2.ResponseImmediateDeliveryDelivery{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/delivery/getall", nil, nil, nil, result)

	return result, err
}

// 拉取已绑定账号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getBindAccount.html
func (comp *Client) GetBindAccount(ctx context.Context) (*response2.ResponseImmediateDeliveryBindAccount, error) {

	result := &response2.ResponseImmediateDeliveryBindAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/local/business/shop/get", nil, nil, nil, result)

	return result, err
}

// 拉取配送单信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getOrder.html
func (comp *Client) GetOrder(ctx context.Context, shopID string, shopOrderID string, shopNO string, deliverySign string) (*response2.ResponseImmediateDeliveryGetOrder, error) {

	result := &response2.ResponseImmediateDeliveryGetOrder{}

	data := &object.StringMap{
		"shopid":        shopID,
		"shop_order_id": shopOrderID,
		"shop_no":       shopNO,
		"delivery_sign": deliverySign,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/addtips", data, nil, nil, result)

	return result, err
}

// 模拟配送公司更新配送单状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.mockUpdateOrder.html
func (comp *Client) MockUpdateOrder(ctx context.Context, options *power.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/test_update_order", options, nil, nil, result)

	return result, err
}

// 第三方代商户发起开通即时配送权限
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.openDelivery.html
func (comp *Client) OpenDelivery(ctx context.Context) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/open", nil, nil, nil, result)

	return result, err
}

// 预下配送单接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preAddOrder.html
func (comp *Client) PreAddOrder(ctx context.Context, data *power.HashMap) (*response2.ResponseImmediateDeliveryPreAddOrder, error) {

	result := &response2.ResponseImmediateDeliveryPreAddOrder{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/pre_add", data, nil, nil, result)

	return result, err
}

// 预取消配送单接口
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preCancelOrder.html
func (comp *Client) PreCancelOrder(ctx context.Context, data *power.HashMap) (*response2.ResponseImmediateDeliveryPreCancelOrder, error) {

	result := &response2.ResponseImmediateDeliveryPreCancelOrder{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/precancel", data, nil, nil, result)

	return result, err
}

// 模拟配送公司更新配送单状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.realMockUpdateOrder.html
func (comp *Client) RealMockUpdateOrder(ctx context.Context, options *power.HashMap) (*response.ResponseMiniProgram, error) {

	result := &response.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/realmock_update_order", options, nil, nil, result)

	return result, err
}

// 重新下单
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.reOrder.html
func (comp *Client) ReOrder(ctx context.Context, data *power.HashMap) (*response2.ResponseImmediateDeliveryReOrder, error) {

	result := &response2.ResponseImmediateDeliveryReOrder{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/immediateDelivery/local/business/order/readd", data, nil, nil, result)

	return result, err
}
