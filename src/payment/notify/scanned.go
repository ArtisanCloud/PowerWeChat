package notify

import (
	"github.com/ArtisanCloud/PowerLibs/v2/http/response"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment/notify/request"
	"net/http"
	"reflect"
)

type Scanned struct {
	*Handler
	alert string
}

func NewScannedNotify(app kernel.ApplicationPaymentInterface, request *http.Request) *Scanned {

	scanned := &Scanned{
		NewHandler(&app, request),
		"",
	}

	scanned.Check = false

	return scanned
}
func (comp *Scanned) Alert(message string) {
	comp.alert = message
}

func (comp *Scanned) Handle(closure func(message *request.RequestNotify, fail func(message string), alert func(message string)) interface{}) (*response.HttpResponse, error) {
	message, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	result := closure(message, comp.Fail, comp.Alert)

	resultCode := FAIL
	if comp.alert == "" && comp.fail == "" {
		resultCode = SUCCESS
	}
	attributes := &object.StringMap{
		"result_code":  resultCode,
		"err_code_des": comp.alert,
	}

	if comp.alert == "" && reflect.TypeOf(result).Name() == "string" {
		config := (*comp.App).GetConfig()
		(*attributes)["appid"] = config.GetString("app_id", "")
		(*attributes)["mch_id"] = config.GetString("mch_id", "")
		(*attributes)["nonce_str"] = object.UniqueID("")
		(*attributes)["prepay_id"] = result.(string)
	}

	return comp.RespondWith(attributes, true).ToResponse()

}
