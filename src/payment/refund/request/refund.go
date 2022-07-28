package request

type RequestRefund struct {
	TransactionID string              `json:"transaction_id,omitempty"`
	OutTradeNo    string              `json:"out_trade_no,omitempty"` // OutTradeNo 和 TransactionID 二选一
	OutRefundNo   string              `json:"out_refund_no,omitempty"`
	Reason        string              `json:"reason,omitempty"`
	NotifyUrl     string              `json:"notify_url,omitempty"`
	FundsAccount  string              `json:"funds_account,omitempty"`
	Amount        *RefundAmount       `json:"amount"`
	GoodsDetail   []*RefundGoodDetail `json:"goods_detail,omitempty"`
}

type RefundAmount struct {
	Refund   int                 `json:"refund"`
	From     []*RefundAmountFrom `json:"from,omitempty"`
	Total    int                 `json:"total"`
	Currency string              `json:"currency"`
}

type RefundAmountFrom struct {
	Account string `json:"account"`
	Amount  int    `json:"amount"`
}

type RefundGoodDetail struct {
	MerchantGoodsID  string `json:"merchant_goods_id"`
	WechatPayGoodsID string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	UnitPrice        int    `json:"unit_price"`
	RefundAmount     int    `json:"refund_amount"`
	RefundQuantity   int    `json:"refund_quantity"`
}
