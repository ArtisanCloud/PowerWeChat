package wifi

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/wifi/request"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/wifi/response"
)

type CardClient struct {
	*kernel.BaseClient
}

func NewCardClient(app kernel.ApplicationInterface) *CardClient {
	return &CardClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// 调用设置门店卡劵投放信息接口
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Cards/Set_store_card_and_coupon_delivery_information.html
func (comp *CardClient) Set(data *request.RequestWifiCardSet) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.HttpPostJson("bizwifi/couponput/set", data, nil, nil, result)

	return result, err

}

// 通过此接口查询某一门店的详细卡券投放信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Cards/Query_the_store_card_and_coupon_delivery_information.html
func (comp *CardClient) Get(shopID int) (*response.ResponseWifiCardGet, error) {

	result := &response.ResponseWifiCardGet{}

	_, err := comp.HttpPostJson("bizwifi/couponput/get", &object.HashMap{"shop_id": shopID}, nil, nil, result)

	return result, err

}
