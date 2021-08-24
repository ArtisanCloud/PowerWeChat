package notify

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/http"
)

type Paid struct {
	*Handler
}

func NewPaidNotify(app kernel.ApplicationPaymentInterface) *Paid {
	paid := &Paid{
		NewHandler(&app),
	}

	return paid
}

func (comp *Paid) Handle(closure func(message *object.HashMap, content *object.HashMap, fail string) interface{}) (*http.Response, error) {

	messages, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	content, err := comp.reqInfo()
	result := closure(messages, content, comp.fail)
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
