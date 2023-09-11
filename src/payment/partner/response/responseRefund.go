package response

import "time"

type Form struct {
	Account string `json:"account"`
	Amount  int    `json:"amount"`
}

type RefundAmount struct {
	Total            int    `json:"total"`
	Refund           int    `json:"refund"`
	From             []Form `json:"from"`
	PayerTotal       int    `json:"payer_total"`
	PayerRefund      int    `json:"payer_refund"`
	SettlementRefund int    `json:"settlement_refund"`
	SettlementTotal  int    `json:"settlement_total"`
	DiscountRefund   int    `json:"discount_refund"`
	Currency         string `json:"currency"`
}

type GoodDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatPayGoodsId string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	UnitPrice        int    `json:"unit_price"`
	RefundAmount     int    `json:"refund_amount"`
	RefundQuantity   int    `json:"refund_quantity"`
}

type PromptDetail struct {
	PromotionId  string       `json:"promotion_id"`
	Scope        string       `json:"scope"`
	Type         string       `json:"type"`
	Amount       int          `json:"amount"`
	RefundAmount int          `json:"refund_amount"`
	GoodsDetail  []GoodDetail `json:"goods_detail"`
}

// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_10.shtml
type ResponseQueryByOutRefundNO struct {
	RefundId            string         `json:"refund_id"`
	OutRefundNo         string         `json:"out_refund_no"`
	TransactionId       string         `json:"transaction_id"`
	OutTradeNo          string         `json:"out_trade_no"`
	Channel             string         `json:"channel"`
	UserReceivedAccount string         `json:"user_received_account"`
	SuccessTime         time.Time      `json:"success_time"`
	CreateTime          time.Time      `json:"create_time"`
	Status              string         `json:"status"`
	FundsAccount        string         `json:"funds_account"`
	Amount              RefundAmount   `json:"amount"`
	PromotionDetail     []PromptDetail `json:"promotion_detail"`
}
