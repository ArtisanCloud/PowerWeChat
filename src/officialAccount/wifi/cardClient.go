package wifi

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/wifi/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/wifi/response"
)

type CardClient struct {
	*kernel.BaseClient
}

func NewCardClient(app kernel.ApplicationInterface) (*CardClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &CardClient{
		baseClient,
	}, nil
}

// 调用设置门店卡劵投放信息接口
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Cards/Set_store_card_and_coupon_delivery_information.html
func (comp *CardClient) Set(ctx context.Context, data *request.RequestWifiCardSet) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/couponput/set", data, nil, nil, result)

	return result, err

}

// 通过此接口查询某一门店的详细卡券投放信息
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Cards/Query_the_store_card_and_coupon_delivery_information.html
func (comp *CardClient) Get(ctx context.Context, shopID int) (*response.ResponseWifiCardGet, error) {

	result := &response.ResponseWifiCardGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "bizwifi/couponput/get", &object.HashMap{"shop_id": shopID}, nil, nil, result)

	return result, err

}
