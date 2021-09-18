package notify

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/http"
)

type Paid struct {
	*Handler
}

func NewPaidNotify(app kernel.ApplicationPaymentInterface, request *http.Request) *Paid {

	paid := &Paid{
		NewHandler(&app, request),
	}

	return paid
}

func (comp *Paid) Handle(closure func(message *power.HashMap, content *power.HashMap, fail func(message string)) interface{}) (*response.HttpResponse, error) {

	hashMessages, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}
	messages, err := power.HashMapToPower(hashMessages)
	if err != nil {
		return nil, err
	}

	hashContent, err := comp.reqInfo()
	content, err := power.HashMapToPower(hashContent)
	if err != nil {
		return nil, err
	}

	result := closure(messages, content, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()

}

func (comp *Paid) reqInfo() (info *object.HashMap, err error) {

	content, err := comp.DecryptMessage("resource")
	if err != nil {
		return nil, err
	}

	info = &object.HashMap{}
	err = object.JsonDecode([]byte(content), info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
