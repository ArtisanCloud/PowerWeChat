package response

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseQueryWorkWX struct {
	*response.ResponsePayment

	XMLName             xml.Name `xml:"xml"`
	Text                string   `xml:",chardata"`
	ReturnCode          string   `xml:"return_code"`
	ReturnMsg           string   `xml:"return_msg"`
	Sign                string   `xml:"sign"`
	ResultCode          string   `xml:"result_code"`
	MchBillno           string   `xml:"mch_billno"`
	MchID               string   `xml:"mch_id"`
	DetailID            string   `xml:"detail_id"`
	Status              string   `xml:"status"`
	SendType            string   `xml:"send_type"`
	TotalAmount         string   `xml:"total_amount"`
	Reason              string   `xml:"reason"`
	SendTime            string   `xml:"send_time"`
	Wishing             string   `xml:"wishing"`
	Remark              string   `xml:"remark"`
	ActName             string   `xml:"act_name"`
	OpenID              string   `xml:"openid"`
	Amount              string   `xml:"amount"`
	RcvTime             string   `xml:"rcv_time"`
	SenderName          string   `xml:"sender_name"`
	SenderHeaderMediaID string   `xml:"sender_header_media_id"`
}
