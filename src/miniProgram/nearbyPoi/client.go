package nearbyPoi

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/nearbyPoi/response"
)

type Client struct {
	*kernel.BaseClient
}

// 添加地点
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html
func (comp *Client) Add(data *power.HashMap) (*response.ResponseNearbyPoiAdd, error) {

	result := &response.ResponseNearbyPoiAdd{}

	_, err := comp.HttpPostJson("wxa/addnearbypoi", data, nil, nil, result)

	return result, err
}

// 删除地点
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html
func (comp *Client) Delete(poiID string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"poi_id": poiID,
	}

	_, err := comp.HttpPostJson("wxa/delnearbypoi", data, nil, nil, result)

	return result, err
}

// 查看地点列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.getList.html
func (comp *Client) GetList(page int, pageRows int) (*response.ResponseNearbyPoiGetList, error) {

	result := &response.ResponseNearbyPoiGetList{}

	params := &object.StringMap{
		"page":      fmt.Sprintf("%d", page),
		"page_rows": fmt.Sprintf("%d", pageRows),
	}

	_, err := comp.HttpGet("wxa/getnearbypoilist", params, nil, result)

	return result, err
}

// 展示/取消展示附近小程序
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.setShowStatus.html
func (comp *Client) SetShowStatus(poiID string, status int) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"poi_id": poiID,
		"status": status,
	}

	_, err := comp.HttpPostJson("wxa/delnearbypoi", data, nil, nil, result)

	return result, err
}
