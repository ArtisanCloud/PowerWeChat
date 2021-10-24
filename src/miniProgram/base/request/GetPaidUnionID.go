package request

type RequestGetPaidUnionID struct {
	OpenID        string `json:"openid"`                   // 支付用户唯一标识
	TransactionID string `json:"transaction_id,omitempty"` //	微信支付订单号
	MchID         string `json:"mch_id,omitempty"`         // 微信支付分配的商户号，和商户订单号配合使用
	OutTradeNo    string `json:"out_trade_no,omitempty"`
}
