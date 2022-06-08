package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseProfitSharingOrder struct {
	response.ResponsePayment

	TransactionID string           `json:"transaction_id"`
	OutOrderNO    string           `json:"out_order_no"`
	OrderID       string           `json:"order_id"`
	State         string           `json:"state"`
	Receivers     []*power.HashMap `json:"receivers"`
}
