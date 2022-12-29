package express

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/express/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 生成运单
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.addOrder.html
func (comp *Client) AddOrder(ctx context.Context, data *power.HashMap) (*response.ResponseExpressAddOrder, error) {

	result := &response.ResponseExpressAddOrder{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/order/add", data, nil, nil, result)

	return result, err
}

// 批量获取运单数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.batchGetOrder.html
func (comp *Client) BatchGetOrder(ctx context.Context, orderList []*power.HashMap) (*response.ResponseExpressBatchOrderList, error) {

	result := &response.ResponseExpressBatchOrderList{}

	data := &object.HashMap{
		"order_list": orderList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/order/batchget", data, nil, nil, result)

	return result, err
}

// 绑定、解绑物流账号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.bindAccount.html
func (comp *Client) BindAccount(ctx context.Context, actionType string, bizID string, deliveryID string, password string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"type":        actionType,
		"biz_id":      bizID,
		"delivery_id": deliveryID,
		"password":    password,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/accountService/bind", data, nil, nil, result)

	return result, err
}

// 取消运单
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.cancelOrder.html
func (comp *Client) CancelOrder(ctx context.Context, orderID string, openID string, deliveryID string, waybillID string) (*response.ResponseExpressCancelOrder, error) {

	result := &response.ResponseExpressCancelOrder{}

	data := &object.HashMap{
		"order_id":    orderID,
		"openid":      openID,
		"delivery_id": deliveryID,
		"waybill_id":  waybillID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/order/cancel", data, nil, nil, result)

	return result, err
}

// 获取所有绑定的物流账号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllAccount.html
func (comp *Client) GetAllAccount(ctx context.Context) (*response.ResponseExpressAccountGetAll, error) {

	result := &response.ResponseExpressAccountGetAll{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/express/business/accountService/getall", nil, nil, result)

	return result, err
}

// 获取支持的快递公司列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllDelivery.html
func (comp *Client) GetAllDelivery(ctx context.Context) (*response.ResponseExpressDeliveryGetAll, error) {

	result := &response.ResponseExpressDeliveryGetAll{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/express/business/delivery/getall", nil, nil, result)

	return result, err
}

// 获取运单数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getOrder.html
func (comp *Client) GetOrder(ctx context.Context, orderID string, openID string, deliveryID string, waybillID string, printType int) (*response.ResponseExpressGetOrder, error) {

	result := &response.ResponseExpressGetOrder{}

	data := &object.HashMap{
		"order_id":    orderID,
		"openid":      openID,
		"delivery_id": deliveryID,
		"waybill_id":  waybillID,
		"print_type":  printType,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/order/get", data, nil, nil, result)

	return result, err
}

// 查询运单轨迹
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPath.html
func (comp *Client) GetPath(ctx context.Context, orderID string, openID string, deliveryID string, waybillID string) (*response.ResponseExpressGetPath, error) {

	result := &response.ResponseExpressGetPath{}

	data := &object.HashMap{
		"order_id":    orderID,
		"openid":      openID,
		"delivery_id": deliveryID,
		"waybill_id":  waybillID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/path/get", data, nil, nil, result)

	return result, err
}

// 获取打印员
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPrinter.html
func (comp *Client) GetPrinter(ctx context.Context) (*response.ResponseExpressGetPrinter, error) {

	result := &response.ResponseExpressGetPrinter{}

	_, err := comp.BaseClient.HttpGet(ctx, "cgi-bin/express/business/path/get", nil, nil, result)

	return result, err
}

// 获取电子面单余额
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getQuota.html
func (comp *Client) GetQuota(ctx context.Context, deliveryID string, bizID string) (*response.ResponseExpressGetQuota, error) {

	result := &response.ResponseExpressGetQuota{}

	data := &object.HashMap{
		"delivery_id": deliveryID,
		"biz_id":      bizID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/quota/get", data, nil, nil, result)

	return result, err
}

// 模拟快递公司更新订单状态, 该接口只能用户测试
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.testUpdateOrder.html
func (comp *Client) TestUpdateOrder(
	ctx context.Context,
	BizID string, OrderID string,
	DeliveryID string, WaybillID string,
	ActionTime int64, ActionType int, ActionMsg string,
) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"biz_id":      BizID,
		"order_id":    OrderID,
		"delivery_id": DeliveryID,
		"waybill_id":  WaybillID,
		"action_time": ActionTime,
		"action_type": ActionType,
		"action_msg":  ActionMsg,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/test_update_order", data, nil, nil, result)

	return result, err
}

// 配置面单打印员
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.updatePrinter.html
func (comp *Client) UpdatePrinter(ctx context.Context, openID string, updateType string, tagidList string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"openid":      openID,
		"update_type": updateType,
		"tagid_list":  tagidList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/business/printer/update", data, nil, nil, result)

	return result, err
}

// 获取面单联系人信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.getContact.html
func (comp *Client) GetContact(ctx context.Context, token string, waybillID string) (*response.ResponseExpressGetContact, error) {

	result := &response.ResponseExpressGetContact{}

	data := &object.HashMap{
		"token":      token,
		"waybill_id": waybillID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/delivery/contact/get", data, nil, nil, result)

	return result, err
}

// 预览面单模板
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.previewTemplate.html
func (comp *Client) PreviewTemplate(ctx context.Context, options *power.HashMap) (*response.ResponseExpressPreviewTemplate, error) {

	result := &response.ResponseExpressPreviewTemplate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/delivery/template/preview", options, nil, nil, result)

	return result, err
}

// 更新商户审核结果
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updateBusiness.html
func (comp *Client) UpdateBusiness(ctx context.Context, options *power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/delivery/service/business/update", options, nil, nil, result)

	return result, err
}

// 更新运单轨迹
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updatePath.html
func (comp *Client) UpdatePath(ctx context.Context, options *power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/express/delivery/path/update?", options, nil, nil, result)

	return result, err
}
