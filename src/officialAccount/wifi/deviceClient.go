package wifi

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/wifi/response"
)

type DeviceClient struct {
	*kernel.BaseClient
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

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Adding_password-protected_devices.html
func (comp *DeviceClient) AddPasswordDevice(ctx *context.Context, shopID int, ssid string, password string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"shop_id":  shopID,
		"ssid":     ssid,
		"password": password,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/device/add", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Adding_password-protected_devices.html
func (comp *DeviceClient) AddPortalDevice(ctx *context.Context, shopID int, ssid string, reset string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"shop_id": shopID,
		"ssid":    ssid,
		"reset":   reset,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/apportal/register", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店删除密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Adding_password-protected_devices.html
func (comp *DeviceClient) Delete(ctx *context.Context, macAddress string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"bssid": macAddress,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/device/delete", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Query_devices.html
func (comp *DeviceClient) List(ctx *context.Context, page int, size int) (*response.ResponseWifiDeviceList, error) {

	result := &response.ResponseWifiDeviceList{}

	params := &object.HashMap{
		"pageindex": page,
		"pagesize":  size,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/device/List", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Query_devices.html
func (comp *DeviceClient) listByShopID(ctx *context.Context, shopID int, page int, size int) (*response.ResponseWifiDeviceList, error) {

	result := &response.ResponseWifiDeviceList{}

	params := &object.HashMap{
		"shop_id":   shopID,
		"pageindex": page,
		"pagesize":  size,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/device/List", params, nil, nil, result)

	return result, err

}
