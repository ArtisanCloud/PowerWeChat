package notify

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/http"
)

type Refund struct {
	*Handler
}

func NewRefundNotify(app kernel.ApplicationPaymentInterface) *Refund {
	paid := &Refund{
		NewHandler(&app),
	}

	return paid
}

func (comp *Refund) Handle(closure func(message *object.HashMap, content *object.HashMap, fail string) interface{}) (*http.Response, error) {

	messages, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	result := closure(messages, nil,comp.fail)
	comp.Strict(result)

	return comp.ToResponse()

}
