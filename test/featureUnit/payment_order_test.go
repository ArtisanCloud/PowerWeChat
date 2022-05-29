package featureUnit

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"testing"
)

func Test_Unify_Order(t *testing.T) {

	response, err := Payment.Order.JSAPITransaction(&power.HashMap{
		"amount": &power.HashMap{
			"total":    1,
			"currency": "CNY",
		},
		"attach":       "自定义数据说明",
		"description":  "Image形象店-深圳腾大-QQ公仔",
		"mchid":        "1611854986",
		"notify_url":   "https://pay.wangchaoyi.com/power/notify",
		"out_trade_no": "5519778939773395659222199361",
		"payer": &power.HashMap{
			"openid": "oAuaP0TRUMwP169nQfg7XCEAw3HQ",
		},
	})

	if err != nil {
		t.Error(err.Error())
	} else if response != nil && response.PrepayID == "" {
		t.Error("response nil")
	}
}
