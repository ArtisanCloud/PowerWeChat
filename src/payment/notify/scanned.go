package notify

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
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

func (comp *Scanned) Handle(closure func(message *power.HashMap, content *power.HashMap, fail string, alert string) interface{}) (*http.Response, error) {
	hashMessages, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}
	messages, err := power.HashMapToPower(hashMessages)
	if err != nil {
		return nil, err
	}

	result := closure(messages, nil, comp.fail, comp.alert)

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
