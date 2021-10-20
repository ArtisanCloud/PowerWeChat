package response

import "encoding/xml"

type ResponseSendWorkWX struct {
	XMLName             xml.Name `xml:"xml"`
	Text                string   `xml:",chardata"`
	ReturnCode          string   `xml:"return_code"`
	ReturnMsg           string   `xml:"return_msg"`
	Sign                string   `xml:"sign"`
	ResultCode          string   `xml:"result_code"`
	MchBillno           string   `xml:"mch_billno"`
	MchID               string   `xml:"mch_id"`
	Wxappid             string   `xml:"wxappid"`
	ReOpenid            string   `xml:"re_openid"`
	TotalAmount         string   `xml:"total_amount"`
	SendListid          string   `xml:"send_listid"`
	SenderName          string   `xml:"sender_name"`
	SenderHeaderMediaID string   `xml:"sender_header_media_id"`
}

