package notify

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/http"
)

type Refund struct {
	*Handler
}

func NewRefundNotify(app kernel.ApplicationPaymentInterface, request *http.Request) *Refund {

	paid := &Refund{
		NewHandler(&app, request),
	}

	return paid
}

func (comp *Refund) Handle(closure func(message *power.HashMap, content *power.HashMap, fail func(message string)) interface{}) (*http.Response, error) {

	hashMessages, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	messages, err := power.HashMapToPower(hashMessages)
	if err != nil {
		return nil, err
	}

	result := closure(messages, nil, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()

}
