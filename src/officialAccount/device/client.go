package device

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/device/request"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/device/response"
)

type Client struct {
	*kernel.BaseClient
}

// 主动发送消息给设备
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-3
func (comp *Client) Message(data *request.RequestDeviceMessage) (*response.ResponseDeviceMessage, error) {

	result := &response.ResponseDeviceMessage{}

	_, err := comp.HttpPostJson("device/transmsg", data, nil, nil, result)

	return result, err
}

// 获取设备二维码
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-4
func (comp *Client) QRCode(deviceIDs []string) (*response.ResponseDeviceCreateQRCode, error) {

	result := &response.ResponseDeviceCreateQRCode{}

	param := &object.HashMap{
		"device_num":     len(deviceIDs),
		"device_id_list": deviceIDs,
	}

	_, err := comp.HttpPostJson("device/create_qrcode", param, nil, nil, result)

	return result, err
}

// 设备授权
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-5
func (comp *Client) Authorize(data request.RequestDeviceAuthorize) (*response.ResponseDeviceAuthorize, error) {

	result := &response.ResponseDeviceAuthorize{}

	_, err := comp.HttpPostJson("device/authorize_device", data, nil, nil, result)

	return result, err
}

// 获取deviceid和二维码
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-6
func (comp *Client) createID(productID string) (*response.ResponseDeviceCreateID, error) {

	result := &response.ResponseDeviceCreateID{}

	params := &object.HashMap{
		"product_id": productID,
	}

	_, err := comp.HttpPostJson("device/authorize_device", params, nil, nil, result)

	return result, err
}

// 设备绑定
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) Bind(openID string, deviceID string, ticket string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"ticket":    ticket,
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.HttpPostJson("device/bind", params, nil, nil, result)

	return result, err
}

// 设备解绑
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) Unbind(openID string, deviceID string, ticket string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"ticket":    ticket,
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.HttpPostJson("device/unbind", params, nil, nil, result)

	return result, err
}



// 设备绑定
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) ForceBind(openID string, deviceID string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.HttpPostJson("device/compel_bind", params, nil, nil, result)

	return result, err
}

// 设备解绑
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (comp *Client) ForceUnbind(openID string, deviceID string) (*response.ResponseDeviceBind, error) {

	result := &response.ResponseDeviceBind{}

	params := &object.HashMap{
		"device_id": deviceID,
		"openid":    openID,
	}

	_, err := comp.HttpPostJson("device/compel_unbind", params, nil, nil, result)

	return result, err
}
