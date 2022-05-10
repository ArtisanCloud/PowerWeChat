package request

type RequestQuery struct {
	TransactionID   string     `json:"transaction_id,omitempty"` // OutTradeNo 和 TransactionID 二选一
	OutOrderNO      string     `json:"out_order_no,omitempty"`
	Receivers       *Receivers `json:"reason,omitempty"`
	UnfreezeUnSplit bool       `json:"notify_url,omitempty"`
}

