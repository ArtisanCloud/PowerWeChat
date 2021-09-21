package notify

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"github.com/ArtisanCloud/power-wechat/src/payment/notify/request"
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

func (comp *Paid) Handle(closure func(request *request.RequestNotify, fail func(message string)) interface{}) (*response.HttpResponse, error) {

	message, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	hashContent, err := comp.reqInfo()
	if err != nil {
		return nil, err
	}

	message.Resource = hashContent

	result := closure(message, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()

}

func (comp *Paid) reqInfo() (info *request.EncryptedResource, err error) {

	content, err := comp.DecryptMessage()
	if err != nil {
		return nil, err
	}

	return info, nil
}
