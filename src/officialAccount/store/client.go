package store

import (
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
func (comp *Client) Categories() (*response.ResponseStoreCategory, error) {

	result := &response.ResponseStoreCategory{}

	_, err := comp.BaseClient.HttpGet("wxa/get_merchant_category", nil, nil, result)

	return result, err
}

// 从腾讯地图拉取省市区信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Districts() (*response.ResponseStoreDistrict, error) {

	result := &response.ResponseStoreDistrict{}

	_, err := comp.BaseClient.HttpGet("wxa/get_district", nil, nil, result)

	return result, err
}

// 在腾讯地图中搜索门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) SearchFromMap(districtID int, keyword string) (*response.ResponseStoreSearchMapPIO, error) {

	result := &response.ResponseStoreSearchMapPIO{}

	params := &object.HashMap{
		"districtid": districtID,
		"keyword":    keyword,
	}

	_, err := comp.BaseClient.HttpPostJson("wxa/search_map_poi", params, nil, nil, result)

	return result, err
}

// 查询门店小程序审核结果
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) GetStatus() (*response.ResponseStoreGetStatus, error) {

	result := &response.ResponseStoreGetStatus{}

	_, err := comp.BaseClient.HttpGet("wxa/get_merchant_audit_info", nil, nil, result)

	return result, err
}

// 修改门店小程序信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) UpdateMerchant(mediaID int, intro string) (*response.ResponseStoreSearchMapPIO, error) {

	result := &response.ResponseStoreSearchMapPIO{}

	params := &object.HashMap{
		"headimg_mediaid": mediaID,
		"intro":           intro,
	}

	_, err := comp.BaseClient.HttpPostJson("wxa/modify_merchant", params, nil, nil, result)

	return result, err
}

// 在腾讯地图中创建门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) CreateFromMap(data *request.BaseInfo) (*response.ResponseStoreCreateFromMap, error) {

	result := &response.ResponseStoreCreateFromMap{}

	_, err := comp.BaseClient.HttpPostJson("wxa/create_map_poi", data, nil, nil, result)

	return result, err
}

// 添加门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Create(data *request.RequestStoreCreate) (*response.ResponseStoreCreate, error) {

	result := &response.ResponseStoreCreate{}

	_, err := comp.BaseClient.HttpPostJson("wxa/add_store", data, nil, nil, result)

	return result, err
}

// 更新门店信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Update(data *request.RequestStoreUpdate) (*response.ResponseStoreUpdate, error) {

	result := &response.ResponseStoreUpdate{}

	_, err := comp.BaseClient.HttpPostJson("wxa/update_store", data, nil, nil, result)

	return result, err
}

// 获取单个门店信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Get(poiID int) (*response.ResponseStoreInfo, error) {

	result := &response.ResponseStoreInfo{}

	params := &object.HashMap{
		"poi_id": poiID,
	}

	_, err := comp.BaseClient.HttpPostJson("wxa/get_store_info", params, nil, nil, result)

	return result, err
}

// 获取门店信息列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) List(offset int, limit int) (*response.ResponseStoreInfo, error) {

	result := &response.ResponseStoreInfo{}

	params := &object.HashMap{
		"offset": offset,
		"limit":  limit,
	}

	_, err := comp.BaseClient.HttpPostJson("wxa/get_store_list", params, nil, nil, result)

	return result, err
}

// 删除门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html
func (comp *Client) Delete(poiID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"poi_id": poiID,
	}

	_, err := comp.BaseClient.HttpPostJson("wxa/del_store", params, nil, nil, result)

	return result, err
}
