package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
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
			"openid": "o4GgauInH_RCEdvrrNGrntXDuXXX",
		},

	}, false)

	if response.PrepayID =="" {
		t.Error("response nil")
	}
	fmt.Dump(err.Error())

	//if response == nil || response.ResponseWork == nil {
	//	t.Error("response nil")
	//} else if response.ErrCode != 0 {
	//	t.Error("response error message as :", response.ErrMSG)
	//}

}
