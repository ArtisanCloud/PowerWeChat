package base

import (
	"github.com/ArtisanCloud/go-libs/object"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
)

type Client struct {
	*kernel.BaseClient
}

// 付款码支付
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1

func (comp *Client) Pay(params *object.StringMap) *response2.ResponseWork {

	result := &response2.ResponseWork{}

	endpoint := comp.Wrap("/v3/pay/micropay")
	comp.Request(endpoint, params, "POST", nil, false, nil, result)

	return result
}

// 付款码查询openid
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9
func (comp *Client) AuthCodeToOpenID(authCode string) *response2.ResponseWork {

	config := (*comp.App).GetConfig()
	appID := config.GetString("app_id", "")

	result := &response2.ResponseWork{}

	comp.Request("tools/authcodetoopenid", &object.StringMap{
		"appid":     appID,
		"auth_code": authCode,
	}, "POST", nil, false, nil, result)

	return result
}
