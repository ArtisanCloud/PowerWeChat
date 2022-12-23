package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml

type ResponseRefund struct {
	response.ResponsePayment

	RefundID            string                   `json:"refund_id"`
	OutRefundNO         string                   `json:"out_refund_no"`
	TransactionID       string                   `json:"transaction_id"`
	OutTradeNO          string                   `json:"out_trade_no"`
	Channel             string                   `json:"channel"`
	UserReceivedAccount string                   `json:"user_received_account"`
	SuccessTime         string                   `json:"success_time"`
	CreateTime          string                   `json:"create_time"`
	Status              string                   `json:"status"`
	FundsAccount        string                   `json:"funds_account"`
	Amount              *RefundAmount            `json:"amount"`
	PromotionDetail     []*RefundPromotionDetail `json:"promotion_detail,omitempty"`
}

type RefundAmount struct {
	Refund int `json:"refund"`
	From   []struct {
		Account string `json:"account"`
		Amount  int    `json:"amount"`
	} `json:"from"`
	Total    int    `json:"total"`
	Currency string `json:"currency"`
}

type GoodsDetail struct {
	MerchantGoodsID  string `json:"merchant_goods_id"`
	WechatPayGoodsID string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	UnitPrice        int    `json:"unit_price"`
	RefundAmount     int    `json:"refund_amount"`
	RefundQuantity   int    `json:"refund_quantity"`
}

type RefundPromotionDetail struct {
	PromotionID  string         `json:"promotion_id"`
	Scope        string         `json:"scope"`
	Type         string         `json:"type"`
	Amount       int            `json:"amount"`
	RefundAmount int            `json:"refund_amount"`
	GoodsDetail  []*GoodsDetail `json:"goods_detail"`
}
