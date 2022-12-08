package response

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_1.shtml

type ResponseProfitSharingReturnOrder struct {
	response.ResponsePayment

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

type ResponseProfitSharingReturn struct {
	XMLName           xml.Name `xml:"xml"`
	Text              string   `xml:",chardata"`
	ReturnCode        string   `xml:"return_code"`
	MchID             string   `xml:"mch_id"`
	Appid             string   `xml:"appid"`
	NonceStr          string   `xml:"nonce_str"`
	Sign              string   `xml:"sign"`
	OrderID           string   `xml:"order_id"`
	OutOrderNo        string   `xml:"out_order_no"`
	OutReturnNo       string   `xml:"out_return_no"`
	ReturnNo          string   `xml:"return_no"`
	ReturnAccountType string   `xml:"return_account_type"`
	ReturnAccount     string   `xml:"return_account"`
	ReturnAmount      string   `xml:"return_amount"`
	Description       string   `xml:"description"`
	Result            string   `xml:"result"`
	FinishTime        string   `xml:"finish_time"`
}
