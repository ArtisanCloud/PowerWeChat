package shakeAround

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround/response"
)

type RelationClient struct {
	*kernel.BaseClient
}

func NewRelationClient(app kernel.ApplicationInterface) *RelationClient {
	return &RelationClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// 新增摇一摇出来的页面信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Pages_management/Page_management.html
func (comp *RelationClient) BindPages(data *request.RequestDeviceIdentifier, pageIDs []int) (*response.ResponsePage, error) {

	result := &response.ResponsePage{}

	params := &object.HashMap{
		"device_identifier": data,
		"page_ids":          pageIDs,
	}
	_, err := comp.HttpPostJson("shakearound/device/bindpage", params, nil, nil, result)

	return result, err
}

// 查询设备与页面的关联关系
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Relationship_between_pages_and_devices/Querying_Device_and_Page_Associations.html
func (comp *RelationClient) ListByDeviceID(data *request.RequestDeviceIdentifier) (*response.ResponseRelationSearch, error) {

	result := &response.ResponseRelationSearch{}

	params := &object.HashMap{
		"type":              1,
		"device_identifier": data,
	}
	_, err := comp.HttpPostJson("shakearound/relation/search", params, nil, nil, result)

	return result, err
}

// 查询设备与页面的关联关系
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Relationship_between_pages_and_devices/Querying_Device_and_Page_Associations.html
func (comp *RelationClient) ListByPageId(pageID int, begin int, count int) (*response.ResponseRelationSearch, error) {

	result := &response.ResponseRelationSearch{}

	params := &object.HashMap{
		"type":    2,
		"page_id": pageID,
		"begin":   begin,
		"count":   count,
	}
	_, err := comp.HttpPostJson("shakearound/relation/search", params, nil, nil, result)

	return result, err
}
