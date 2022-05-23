package wifi

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/wifi/response"
)

type DeviceClient struct {
	*kernel.BaseClient
}


func NewDeviceClient(app kernel.ApplicationInterface) *DeviceClient {
	return &DeviceClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Adding_password-protected_devices.html
func (comp *DeviceClient) AddPasswordDevice(shopID int, ssid string, password string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"shop_id":  shopID,
		"ssid":     ssid,
		"password": password,
	}

	_, err := comp.HttpPostJson("bizwifi/device/add", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Adding_password-protected_devices.html
func (comp *DeviceClient) AddPortalDevice(shopID int, ssid string, reset string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"shop_id": shopID,
		"ssid":    ssid,
		"reset":   reset,
	}

	_, err := comp.HttpPostJson("bizwifi/apportal/register", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店删除密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Adding_password-protected_devices.html
func (comp *DeviceClient) Delete(macAddress string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"bssid": macAddress,
	}

	_, err := comp.HttpPostJson("bizwifi/device/delete", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Query_devices.html
func (comp *DeviceClient) List(page int, size int) (*response.ResponseWifiDeviceList, error) {

	result := &response.ResponseWifiDeviceList{}

	params := &object.HashMap{
		"pageindex": page,
		"pagesize":  size,
	}

	_, err := comp.HttpPostJson("bizwifi/device/List", params, nil, nil, result)

	return result, err

}

// 调用此接口向指定门店添加密码型设备的 Wi-Fi 信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Devices_Management/Query_devices.html
func (comp *DeviceClient) listByShopID(shopID int, page int, size int) (*response.ResponseWifiDeviceList, error) {

	result := &response.ResponseWifiDeviceList{}

	params := &object.HashMap{
		"shop_id":   shopID,
		"pageindex": page,
		"pagesize":  size,
	}

	_, err := comp.HttpPostJson("bizwifi/device/List", params, nil, nil, result)

	return result, err

}
