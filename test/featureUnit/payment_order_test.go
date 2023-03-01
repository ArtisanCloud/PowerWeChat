package featureUnit

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order/request"
	"testing"
)

func Test_Unify_Order(t *testing.T) {

	response, err := Payment.Order.JSAPITransaction(nil, &request.RequestJSAPIPrepay{
		Amount: &request.JSAPIAmount{
			Total:    1,
			Currency: "CNY",
		},
		Attach:      "自定义数据说明",
		Description: "Image形象店-深圳腾大-QQ公仔",
		OutTradeNo:  "5519778939773395659222199361",
		Payer: &request.JSAPIPayer{
			OpenID: "oAuaP0TRUMwP169nQfg7XCEAw3HQ",
		},
	})

	if err != nil {
		t.Error(err.Error())
	} else if response != nil && response.PrepayID == "" {
		t.Error("response nil")
	}
}
