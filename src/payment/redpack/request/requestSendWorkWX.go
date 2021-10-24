package request

import "encoding/xml"

type RequestSendWorkWX struct {
	XMLName             xml.Name `xml:"xml"`
	Text                string   `xml:",chardata"`
	NonceStr            string   `xml:"nonce_str"`
	Sign                string   `xml:"sign"`
	MchBillno           string   `xml:"mch_billno"`
	MchID               string   `xml:"mch_id"`
	Wxappid             string   `xml:"wxappid"`
	SenderName          string   `xml:"sender_name"`
	SenderHeaderMediaID string   `xml:"sender_header_media_id"`
	ReOpenid            string   `xml:"re_openid"`
	TotalAmount         string   `xml:"total_amount"`
	Wishing             string   `xml:"wishing"`
	ActName             string   `xml:"act_name"`
	Remark              string   `xml:"remark"`
	WorkwxSign          string   `xml:"workwx_sign"`
}

