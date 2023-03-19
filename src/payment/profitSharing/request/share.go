package request

import (
	"encoding/xml"
	"time"
)

type RequestShare struct {
	AppID           string      `json:"appid,omitempty"`
	TransactionID   string      `json:"transaction_id,omitempty"` // OutTradeNo 和 TransactionID 二选一
	OutOrderNO      string      `json:"out_order_no,omitempty"`
	Receivers       []*Receiver `json:"receivers,omitempty"`
	UnfreezeUnSplit bool        `json:"unfreeze_unsplit,omitempty"`
}

type Receiver struct {
	Type        string `json:"type"`
	Account     string `json:"account"`
	Name        string `json:"name,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}

type RequestShareReturn struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`

	AppID string `xml:"appid"`
	MchID string `xml:"mch_id"`
	//NonceStr string `xml:"nonce_str"`
	//SignType          string `xml:"sign_type"`
	//Sign              string `xml:"sign"`
	OutOrderNo        string `xml:"out_order_no"`
	OutReturnNo       string `xml:"out_return_no"`
	ReturnAccountType string `xml:"return_account_type"`
	ReturnAccount     string `xml:"return_account"`
	ReturnAmount      string `xml:"return_amount"`
	Description       string `xml:"description"`
}

type ReceiverShareResult struct {
	Amount      int64     `json:"amount,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	Account     string    `json:"account"`
	Result      string    `json:"result"`
	FailReason  string    `json:"fail_reason,omitempty"`
	DetailId    string    `json:"detail_id,omitempty"`
	CreateTime  time.Time `json:"create_time"`
	FinishTime  time.Time `json:"finish_time"`
}
