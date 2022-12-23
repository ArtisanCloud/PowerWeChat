package notify

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/notify/request"
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

func (comp *Paid) Handle(closure func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{}) (*http.Response, error) {

	message, err := comp.GetMessage()
	if err != nil {
		return nil, err
	}

	reqInfo, err := comp.reqInfo()
	if err != nil {
		return nil, err
	}

	// struct the content
	transaction := &models.Transaction{}
	err = object.JsonDecode([]byte(reqInfo), transaction)
	if err != nil {
		return nil, err
	}

	result := closure(message, transaction, comp.Fail)
	comp.Strict(result)

	return comp.ToResponse()

}
