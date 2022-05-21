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
func (comp *DeviceClient) Apply(data *request.RequestDeviceApply) (*response.ResponseDeviceApply, error) {

	result := &response.ResponseDeviceApply{}

	_, err := comp.HttpPostJson("shakearound/device/applyid", data, nil, nil, result)

	return result, err
}

// 查询设备 ID 申请的审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_the_audit_status_of_the_device_ID.html
func (comp *DeviceClient) Status(data *request.RequestDeviceApplyStatus) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	_, err := comp.HttpPostJson("shakearound/device/applystatus", data, nil, nil, result)

	return result, err
}

// 查询设备 ID 申请的审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_the_audit_status_of_the_device_ID.html
func (comp *DeviceClient) DeviceUpdate(deviceIdentifier *request.RequestDeviceIdentifier, comment string) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &object.HashMap{
		"device_identifier": deviceIdentifier,
		"comment":           comment,
	}

	_, err := comp.HttpPostJson("shakearound/device/update", params, nil, nil, result)

	return result, err
}

// 直接关联在设备
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Configure_the_connected_relationship_between_the_device_and_the_store.html
func (comp *DeviceClient) BindPoi(deviceIdentifier *request.RequestDeviceIdentifier, poiID int) (*response.ResponseDeviceBindPoi, error) {

	result := &response.ResponseDeviceBindPoi{}

	params := &object.HashMap{
		"device_identifier": deviceIdentifier,
		"poi_id":            poiID,
	}

	_, err := comp.HttpPostJson("shakearound/device/bindlocation", params, nil, nil, result)

	return result, err
}

// 直接关联在第三方设备
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Configure_the_connected_relationship_between_the_device_and_the_store.html
func (comp *DeviceClient) BindThirdPoi(deviceIdentifier *request.RequestDeviceIdentifier, poiID int, appID string) (*response.ResponseDeviceBindPoi, error) {

	result := &response.ResponseDeviceBindPoi{}

	params := &object.HashMap{
		"device_identifier": deviceIdentifier,
		"poi_id":            poiID,
		"type":              2,
		"poi_appid":         appID,
	}

	_, err := comp.HttpPostJson("shakearound/device/bindlocation", params, nil, nil, result)

	return result, err
}

// 查询已有的设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_device_list.html
func (comp *DeviceClient) ListByIDs(deviceIdentifier []*request.RequestDeviceIdentifier) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &request.RequestDeviceSearch{
		Type:              1,
		DeviceIdentifiers: deviceIdentifier,
	}

	err := comp.search(params, result)

	return result, err
}

// 查询已有的设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_device_list.html
func (comp *DeviceClient) List(lastID int, count int) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &request.RequestDeviceSearch{
		Type:     2,
		LastSeen: lastID,
		Count:    count,
	}

	err := comp.search(params, result)

	return result, err
}

// 查询已有的设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_device_list.html
func (comp *DeviceClient) ListByApplyID(applyID int, lastID int, count int) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &request.RequestDeviceSearch{
		Type:     3,
		ApplyID:  applyID,
		LastSeen: lastID,
		Count:    count,
	}

	err := comp.search(params, result)

	return result, err
}

func (comp *DeviceClient) search(data *request.RequestDeviceSearch, result *response.ResponseDeviceApplyStatus) error {

	_, err := comp.HttpPostJson("shakearound/device/search", data, nil, nil, result)

	return err
}
