package shakeAround

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/shakeAround/response"
)

type GroupClient struct {
	*kernel.BaseClient
}

func NewGroupClient(app kernel.ApplicationInterface) (*GroupClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &GroupClient{
		baseClient,
	}, nil
}

// 新建设备分组
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Active_from_Html5/New_groups.html
func (comp *GroupClient) Create(name string) (*response.ResponseGroup, error) {

	result := &response.ResponseGroup{}

	params := &object.HashMap{
		"group_name": name,
	}

	_, err := comp.HttpPostJson("shakearound/device/group/add", params, nil, nil, result)

	return result, err
}

// 更新设备分组
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Active_from_Html5/New_groups.html
func (comp *GroupClient) Update(groupID string, name string) (*response.ResponseGroup, error) {

	result := &response.ResponseGroup{}

	params := &object.HashMap{
		"group_id":   groupID,
		"group_name": name,
	}

	_, err := comp.HttpPostJson("shakearound/device/group/update", params, nil, nil, result)

	return result, err
}

// 删除设备分组
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Active_from_Html5/New_groups.html
func (comp *GroupClient) Delete(groupID string) (*response.ResponseGroup, error) {

	result := &response.ResponseGroup{}

	params := &object.HashMap{
		"group_id": groupID,
	}

	_, err := comp.HttpPostJson("shakearound/device/group/delete", params, nil, nil, result)

	return result, err
}

// 更新设备分组
// https://developers.weixin.qq.com/doc/offiaccount/en/Shake_Nearby/Active_from_Html5/Search_groups_list.html
func (comp *GroupClient) List(begin int, count int) (*response.ResponseGroupList, error) {

	result := &response.ResponseGroupList{}

	params := &object.HashMap{
		"begin": begin,
		"count": count,
	}

	_, err := comp.HttpPostJson("shakearound/device/group/getlist", params, nil, nil, result)

	return result, err
}

// 查询分组详情
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Active_from_Html5/Search_grouping_details.html
func (comp *GroupClient) Get(groupID int, begin int, count int) (*response.ResponseGroupDetail, error) {

	result := &response.ResponseGroupDetail{}

	params := &object.HashMap{
		"group_id": groupID,
		"begin":    begin,
		"count":    count,
	}

	_, err := comp.HttpPostJson("shakearound/device/group/getdetail", params, nil, nil, result)

	return result, err
}

// 添加设备到分组
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Active_from_Html5/Add_device_to_group.html
func (comp *GroupClient) AddDevice(groupID int, deviceIdentifiers []*request.RequestDeviceIdentifier) (*response.ResponseGroup, error) {

	result := &response.ResponseGroup{}

	params := &object.HashMap{
		"group_id":           groupID,
		"device_identifiers": deviceIdentifiers,
	}

	_, err := comp.HttpPostJson("shakearound/device/group/adddevice", params, nil, nil, result)

	return result, err
}

// 从分组中移除设备
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Active_from_Html5/Remove_device_from_group.html
func (comp *GroupClient) RemoveDevices(groupID int, deviceIdentifiers []*request.RequestDeviceIdentifier) (*response.ResponseGroup, error) {

	result := &response.ResponseGroup{}

	params := &object.HashMap{
		"group_id":           groupID,
		"device_identifiers": deviceIdentifiers,
	}

	_, err := comp.HttpPostJson("shakearound/device/group/deletedevice", params, nil, nil, result)

	return result, err
}
