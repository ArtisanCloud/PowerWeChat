package request

type RequestShare struct {
	AppID           string     `json:"appid,omitempty"`
	TransactionID   string     `json:"transaction_id,omitempty"` // OutTradeNo 和 TransactionID 二选一
	OutOrderNO      string     `json:"out_order_no,omitempty"`
	Receivers       *Receivers `json:"reason,omitempty"`
	UnfreezeUnSplit bool       `json:"notify_url,omitempty"`
}

type Receivers struct {
	Type        string `json:"type"`
	Account     string `json:"account"`
	Name        string `json:"name"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}
