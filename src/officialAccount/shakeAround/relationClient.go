package shakeAround

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/response"
)

type RelationClient struct {
	BaseClient *kernel.BaseClient
}

func NewRelationClient(app kernel.ApplicationInterface) (*RelationClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &RelationClient{
		baseClient,
	}, nil
}

// 新增摇一摇出来的页面信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Pages_management/Page_management.html
func (comp *RelationClient) BindPages(data *request.RequestDeviceIdentifier, pageIDs []int) (*response.ResponsePage, error) {

	result := &response.ResponsePage{}

	params := &object.HashMap{
		"device_identifier": data,
		"page_ids":          pageIDs,
	}
	_, err := comp.BaseClient.HttpPostJson("shakearound/device/bindpage", params, nil, nil, result)

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
	_, err := comp.BaseClient.HttpPostJson("shakearound/relation/search", params, nil, nil, result)

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
	_, err := comp.BaseClient.HttpPostJson("shakearound/relation/search", params, nil, nil, result)

	return result, err
}
