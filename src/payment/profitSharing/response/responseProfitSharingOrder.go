package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
)

type ResponseProfitSharingOrder struct {
	TransactionID string           `json:"transaction_id"`
	OutOrderNO    string           `json:"out_order_no"`
	OrderID       string           `json:"order_id"`
	State         string           `json:"state"`
	Receivers     []*power.HashMap `json:"receivers"`
}
