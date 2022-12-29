package shakeAround

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/response"
)

type DeviceClient struct {
	BaseClient *kernel.BaseClient
}

func NewDeviceClient(app kernel.ApplicationInterface) (*DeviceClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &DeviceClient{
		baseClient,
	}, nil
}

// 申请配置设备
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Apply_device_ID.html
func (comp *DeviceClient) Apply(ctx *context.Context, data *request.RequestDeviceApply) (*response.ResponseDeviceApply, error) {

	result := &response.ResponseDeviceApply{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/device/applyid", data, nil, nil, result)

	return result, err
}

// 查询设备 ID 申请的审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_the_audit_status_of_the_device_ID.html
func (comp *DeviceClient) Status(ctx *context.Context, data *request.RequestDeviceApplyStatus) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/device/applystatus", data, nil, nil, result)

	return result, err
}

// 查询设备 ID 申请的审核状态
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_the_audit_status_of_the_device_ID.html
func (comp *DeviceClient) DeviceUpdate(ctx *context.Context, deviceIdentifier *request.RequestDeviceIdentifier, comment string) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &object.HashMap{
		"device_identifier": deviceIdentifier,
		"comment":           comment,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/device/update", params, nil, nil, result)

	return result, err
}

// 直接关联在设备
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Configure_the_connected_relationship_between_the_device_and_the_store.html
func (comp *DeviceClient) BindPoi(ctx *context.Context, deviceIdentifier *request.RequestDeviceIdentifier, poiID int) (*response.ResponseDeviceBindPoi, error) {

	result := &response.ResponseDeviceBindPoi{}

	params := &object.HashMap{
		"device_identifier": deviceIdentifier,
		"poi_id":            poiID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/device/bindlocation", params, nil, nil, result)

	return result, err
}

// 直接关联在第三方设备
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Configure_the_connected_relationship_between_the_device_and_the_store.html
func (comp *DeviceClient) BindThirdPoi(ctx *context.Context, deviceIdentifier *request.RequestDeviceIdentifier, poiID int, appID string) (*response.ResponseDeviceBindPoi, error) {

	result := &response.ResponseDeviceBindPoi{}

	params := &object.HashMap{
		"device_identifier": deviceIdentifier,
		"poi_id":            poiID,
		"type":              2,
		"poi_appid":         appID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/device/bindlocation", params, nil, nil, result)

	return result, err
}

// 查询已有的设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_device_list.html
func (comp *DeviceClient) ListByIDs(ctx *context.Context, deviceIdentifier []*request.RequestDeviceIdentifier) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &request.RequestDeviceSearch{
		Type:              1,
		DeviceIdentifiers: deviceIdentifier,
	}

	err := comp.search(ctx, params, result)

	return result, err
}

// 查询已有的设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_device_list.html
func (comp *DeviceClient) List(ctx *context.Context, lastID int, count int) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &request.RequestDeviceSearch{
		Type:     2,
		LastSeen: lastID,
		Count:    count,
	}

	err := comp.search(ctx, params, result)

	return result, err
}

// 查询已有的设备信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Devices_management/Query_device_list.html
func (comp *DeviceClient) ListByApplyID(ctx *context.Context, applyID int, lastID int, count int) (*response.ResponseDeviceApplyStatus, error) {

	result := &response.ResponseDeviceApplyStatus{}

	params := &request.RequestDeviceSearch{
		Type:     3,
		ApplyID:  applyID,
		LastSeen: lastID,
		Count:    count,
	}

	err := comp.search(ctx, params, result)

	return result, err
}

func (comp *DeviceClient) search(ctx *context.Context, data *request.RequestDeviceSearch, result *response.ResponseDeviceApplyStatus) error {

	_, err := comp.BaseClient.HttpPostJson(ctx, "shakearound/device/search", data, nil, nil, result)

	return err
}
