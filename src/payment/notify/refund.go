package notify

import (
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/http"
)

type Refund struct {
	*Handler
}

func NewRefund(app kernel.ApplicationPaymentInterface) *Refund {
	paid := &Refund{
		NewHandler(&app),
	}

	return paid
}

func (comp *Refund) Handle(closure func(payload ...interface{}) interface{}) *http.Response {

	result := closure(comp.GetMessage(), comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()

}
