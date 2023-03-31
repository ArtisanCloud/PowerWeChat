package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseImmediateDeliveryPreAddOrder struct {
	response.ResponseMiniProgram

	Fee              int     `json:"fee"`               // 10,
	DeliverFee       string  `json:"deliverfee"`        // 10,
	CouponFee        string  `json:"couponfee"`         // 0,
	Tips             string  `json:"tips"`              // 0,
	InsurancFee      float64 `json:"insurancfee"`       // 0,
	Distance         float64 `json:"distance"`          // 1000,
	DispatchDuration int64   `json:"dispatch_duration"` // 300,
	DeliveryToken    int64   `json:"delivery_token"`    //" 1111111"

}
