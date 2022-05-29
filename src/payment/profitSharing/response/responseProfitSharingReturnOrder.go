package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_1.shtml

type ResponseProfitSharingReturnOrder struct {
	*response.ResponsePayment

	OrderID     string  `json:"order_id"`
	OutOrderNO  string  `json:"out_order_no"`
	OutReturnNO string  `json:"out_return_no"`
	ReturnID    string  `json:"return_id"`
	ReturnMchID string  `json:"return_mchid"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Result      string  `json:"result"`
	FailReason  string  `json:"fail_reason"`
	CreateTime  string  `json:"create_time"`
	FinishTime  string  `json:"finish_time"`
}
