package store

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/store/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/store/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 拉取门店小程序类目
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Categories(ctx *context.Context) (*response.ResponseStoreCategory, error) {

	result := &response.ResponseStoreCategory{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/get_merchant_category", nil, nil, result)

	return result, err
}

// 从腾讯地图拉取省市区信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Districts(ctx *context.Context) (*response.ResponseStoreDistrict, error) {

	result := &response.ResponseStoreDistrict{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/get_district", nil, nil, result)

	return result, err
}

// 在腾讯地图中搜索门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) SearchFromMap(ctx *context.Context, districtID int, keyword string) (*response.ResponseStoreSearchMapPIO, error) {

	result := &response.ResponseStoreSearchMapPIO{}

	params := &object.HashMap{
		"districtid": districtID,
		"keyword":    keyword,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/search_map_poi", params, nil, nil, result)

	return result, err
}

// 查询门店小程序审核结果
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) GetStatus(ctx *context.Context) (*response.ResponseStoreGetStatus, error) {

	result := &response.ResponseStoreGetStatus{}

	_, err := comp.BaseClient.HttpGet(ctx, "wxa/get_merchant_audit_info", nil, nil, result)

	return result, err
}

// 修改门店小程序信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) UpdateMerchant(ctx *context.Context, mediaID int, intro string) (*response.ResponseStoreSearchMapPIO, error) {

	result := &response.ResponseStoreSearchMapPIO{}

	params := &object.HashMap{
		"headimg_mediaid": mediaID,
		"intro":           intro,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/modify_merchant", params, nil, nil, result)

	return result, err
}

// 在腾讯地图中创建门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) CreateFromMap(ctx *context.Context, data *request.BaseInfo) (*response.ResponseStoreCreateFromMap, error) {

	result := &response.ResponseStoreCreateFromMap{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/create_map_poi", data, nil, nil, result)

	return result, err
}

// 添加门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Create(ctx *context.Context, data *request.RequestStoreCreate) (*response.ResponseStoreCreate, error) {

	result := &response.ResponseStoreCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/add_store", data, nil, nil, result)

	return result, err
}

// 更新门店信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Update(ctx *context.Context, data *request.RequestStoreUpdate) (*response.ResponseStoreUpdate, error) {

	result := &response.ResponseStoreUpdate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/update_store", data, nil, nil, result)

	return result, err
}

// 获取单个门店信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Get(ctx *context.Context, poiID int) (*response.ResponseStoreInfo, error) {

	result := &response.ResponseStoreInfo{}

	params := &object.HashMap{
		"poi_id": poiID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/get_store_info", params, nil, nil, result)

	return result, err
}

// 获取门店信息列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) List(ctx *context.Context, offset int, limit int) (*response.ResponseStoreInfo, error) {

	result := &response.ResponseStoreInfo{}

	params := &object.HashMap{
		"offset": offset,
		"limit":  limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/get_store_list", params, nil, nil, result)

	return result, err
}

// 删除门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Delete(ctx *context.Context, poiID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"poi_id": poiID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/del_store", params, nil, nil, result)

	return result, err
}
