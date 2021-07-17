package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/object"
	"testing"
)

func Test_Unify_Order(t *testing.T) {

	response := Payment.Order.Unify(&object.StringMap{
		"out_trade_no": "order-no-for-dev",
	}, false)

	if response.PrepayID =="" {
		t.Error("response nil")
	}

	//if response == nil || response.ResponseWork == nil {
	//	t.Error("response nil")
	//} else if response.ErrCode != 0 {
	//	t.Error("response error message as :", response.ErrMSG)
	//}

}
