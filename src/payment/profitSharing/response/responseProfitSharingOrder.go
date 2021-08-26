package response

import "github.com/ArtisanCloud/go-libs/object"

type ResponseProfitSharingOrder struct {
	TransactionID string            `json:"transaction_id"`
	OutOrderNO    string            `json:"out_order_no"`
	OrderID       string            `json:"order_id"`
	State         string            `json:"state"`
	Receivers     []*object.HashMap `json:"receivers"`
}
