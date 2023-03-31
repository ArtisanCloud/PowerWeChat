package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseImmediateDeliveryReOrder struct {
	response.ResponseMiniProgram

	Fee              int     `json:"fee"`               //  10,
	DeliverFee       string  `json:"deliverfee"`        //  10,
	CouponFee        string  `json:"couponfee"`         //  0,
	Tips             string  `json:"tips"`              //  0,
	InsuranceFee     float64 `json:"insurancfee"`       //  0,
	Distance         float64 `json:"distance"`          //  1000,
	WaybillID        int64   `json:"waybill_id"`        //  "123456789",
	OrderStatus      int64   `json:"order_status"`      //  101,
	FinishCode       int64   `json:"finish_code"`       //  1024,
	PickupCode       int64   `json:"pickup_code"`       //  2048,
	DispatchDuration int64   `json:"dispatch_duration"` //  300

}
