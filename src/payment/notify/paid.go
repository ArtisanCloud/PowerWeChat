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

func (comp *Paid) Handle(closure func(payload ...interface{}) interface{}) (*http.Response, error) {

	messages, err := comp.GetMessage()

	result := closure(messages, comp.reqInfo, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse(), err

}

func (comp *Paid) reqInfo() *object.HashMap {
	return nil
}
