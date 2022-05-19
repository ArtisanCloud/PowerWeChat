package response

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseSendWorkWX struct {
	*response.ResponsePayment

	XMLName             xml.Name `xml:"xml"`
	Text                string   `xml:",chardata"`
	ReturnCode          string   `xml:"return_code"`
	ReturnMsg           string   `xml:"return_msg"`
	Sign                string   `xml:"sign"`
	ResultCode          string   `xml:"result_code"`
	MchBillNO           string   `xml:"mch_billno"`
	MchID               string   `xml:"mch_id"`
	WXAppID             string   `xml:"wxappid"`
	ReOpenID            string   `xml:"re_openid"`
	TotalAmount         string   `xml:"total_amount"`
	SendListID          string   `xml:"send_listid"`
	SenderName          string   `xml:"sender_name"`
	SenderHeaderMediaID string   `xml:"sender_header_media_id"`
}

