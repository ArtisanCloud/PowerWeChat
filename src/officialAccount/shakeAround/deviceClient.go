package shakeAround

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround/response"
)

type DeviceClient struct {
	*kernel.BaseClient
}

func NewDeviceClient(app kernel.ApplicationInterface) *DeviceClient {
	return &DeviceClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// 申请配置设备
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Apply_device_ID.html
func (comp *Client) Apply(data *request.RequestDeviceApply) (*response.ResponseDeviceApply, error) {

	result := &response.ResponseDeviceApply{}

	_, err := comp.HttpPostJson("shakearound/device/applyid", data, nil, nil, result)

	return result, err
}

// 查询设备 ID 申请的审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_the_audit_status_of_the_device_ID.html
func (comp *Client) ApplyStatus(data *request.RequestDeviceApplyStatus) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	_, err := comp.HttpPostJson("shakearound/device/applystatus", data, nil, nil, result)

	return result, err
}

// 查询设备 ID 申请的审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_the_audit_status_of_the_device_ID.html
func (comp *Client) DeviceUpdate(deviceIdentifier *request.RequestDeviceUpdate, comment string) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &object.HashMap{
		"device_identifier": deviceIdentifier,
		"comment":           comment,
	}

	_, err := comp.HttpPostJson("shakearound/device/update", params, nil, nil, result)

	return result, err
}
