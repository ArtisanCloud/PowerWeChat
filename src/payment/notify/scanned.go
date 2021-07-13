package notify

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/power-wechat/src/payment"
	"net/http"
	"reflect"
)

type Scanned struct {
	*Handler
	alert string
}

func NewScanned(app *payment.Payment) *Scanned {
	scanned := &Scanned{
		NewHandler(app),
		"",
	}

	scanned.Check = false

	return scanned
}
func (comp *Scanned) Alert(message string) {
	comp.alert = message
}

func (comp *Scanned) Handle(closure func(payload ...interface{}) interface{}) *http.Response {

	result := closure(comp.GetMessage(), comp.Fail, comp.Alert)

	resultCode := FAIL
	if comp.alert == "" && comp.Fail == "" {
		resultCode = SUCCESS
	}
	attributes := &object.HashMap{
		"result_code":  resultCode,
		"err_code_des": comp.alert,
	}

	if comp.alert == "" && reflect.TypeOf(result).Name() == "string" {
		(*attributes)["appid"] = comp.App.Config.GetString("app_id", "")
		(*attributes)["mch_id"] = comp.App.Config.GetString("mch_id", "")
		(*attributes)["nonce_str"] = str.UniqueID("")
		(*attributes)["prepay_id"] = result
	}

	return comp.RespondWith(attributes, true).ToResponse()

}
