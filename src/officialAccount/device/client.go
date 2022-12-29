package device

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/device/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/device/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 主动发送消息给设备
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-3
func (comp *Client) Message(ctx *context.Context, data *request.RequestDeviceMessage) (*response.ResponseDeviceMessage, error) {

	result := &response.ResponseDeviceMessage{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/transmsg", data, nil, nil, result)

	return result, err
}

// 获取设备二维码
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-4
func (comp *Client) QRCode(ctx *context.Context, deviceIDs []string) (*response.ResponseDeviceCreateQRCode, error) {

	result := &response.ResponseDeviceCreateQRCode{}

	param := &object.HashMap{
		"device_num":     len(deviceIDs),
		"device_id_list": deviceIDs,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/create_qrcode", param, nil, nil, result)

	return result, err
}

// 设备授权
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-5
func (comp *Client) Authorize(ctx *context.Context, data request.RequestDeviceAuthorize) (*response.ResponseDeviceAuthorize, error) {

	result := &response.ResponseDeviceAuthorize{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/authorize_device", data, nil, nil, result)

	return result, err
}

// 获取deviceid和二维码
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-6
func (comp *Client) createID(ctx *context.Context, productID string) (*response.ResponseDeviceCreateID, error) {

	result := &response.ResponseDeviceCreateID{}

	params := &object.HashMap{
		"product_id": productID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/authorize_device", params, nil, nil, result)

	return result, err
}

// 设备绑定
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) Bind(ctx *context.Context, openID string, deviceID string, ticket string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"ticket":    ticket,
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/bind", params, nil, nil, result)

	return result, err
}

// 设备解绑
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) Unbind(ctx *context.Context, openID string, deviceID string, ticket string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"ticket":    ticket,
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/unbind", params, nil, nil, result)

	return result, err
}

// 设备绑定
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) ForceBind(ctx *context.Context, openID string, deviceID string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/compel_bind", params, nil, nil, result)

	return result, err
}

// 设备解绑
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) ForceUnbind(ctx *context.Context, openID string, deviceID string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "device/compel_unbind", params, nil, nil, result)

	return result, err
}
