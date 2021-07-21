package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/object"
	"testing"
)

func Test_Unify_Order(t *testing.T) {

	response ,err:= Payment.Order.Unify(&object.HashMap{
		"out_trade_no": "1217752501201407033233368318",
		"description": "保险机构请款",
		"amount": &object.HashMap{
			"total": 1,
			"currency": "CNY",
		},
		"payer": &object.HashMap{
			"openid": "oAuaP0TRUMwP169nQfg7XCEAw3HQ",
		},

	}, false)

	//response ,err:= Payment.Order.Unify(&object.HashMap{
	//	"amount": &object.HashMap{
	//		"total": 1,
	//		"currency": "CNY",
	//	},
	//	"attach": "自定义数据说明",
	//	"description": "Image形象店-深圳腾大-QQ公仔",
	//	"mchid": "1611854986",
	//	"notify_url": "https://pay.wangchaoyi.com/wx/notify",
	//	"out_trade_no": "5519778939773395659222199361",
	//	"payer": &object.HashMap{
	//		"openid": "oAuaP0TRUMwP169nQfg7XCEAw3HQ",
	//	},
	//}, false)


	if err!=nil{
		t.Error(err.Error())
	}else if response!=nil && response.PrepayID =="" {
		t.Error("response nil")
	}

	//if response == nil || response.ResponseWork == nil {
	//	t.Error("response nil")
	//} else if response.ErrCode != 0 {
	//	t.Error("response error message as :", response.ErrMSG)
	//}

}
