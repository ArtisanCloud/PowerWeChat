package nearbyPoi

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/nearbyPoi/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 添加地点
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html
func (comp *Client) Add(ctx context.Context, data *power.HashMap) (*response.ResponseNearbyPoiAdd, error) {

	result := &response.ResponseNearbyPoiAdd{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/addnearbypoi", data, nil, nil, result)

	return result, err
}

// 删除地点
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html
func (comp *Client) Delete(ctx context.Context, poiID string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"poi_id": poiID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/delnearbypoi", data, nil, nil, result)

	return result, err
}

// 查看地点列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.getList.html
func (comp *Client) GetList(ctx context.Context, page int, pageRows int) (*response.ResponseNearbyPoiGetList, error) {

	result := &response.ResponseNearbyPoiGetList{}

	params := &object.StringMap{
		"page":      fmt.Sprintf("%d", page),
		"page_rows": fmt.Sprintf("%d", pageRows),
	}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/getnearbypoilist", params, nil, result)

	return result, err
}

// 展示/取消展示附近小程序
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.setShowStatus.html
func (comp *Client) SetShowStatus(ctx context.Context, poiID string, status int) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"poi_id": poiID,
		"status": status,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/delnearbypoi", data, nil, nil, result)

	return result, err
}
