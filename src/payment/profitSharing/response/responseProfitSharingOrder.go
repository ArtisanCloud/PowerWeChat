package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/profitSharing/request"
)

type ResponseProfitSharingOrder struct {
	response.ResponsePayment

	TransactionID string                         `json:"transaction_id"`
	OutOrderNO    string                         `json:"out_order_no"`
	OrderID       string                         `json:"order_id"`
	State         string                         `json:"state"`
	Receivers     []*request.ReceiverShareResult `json:"receivers"`
}
