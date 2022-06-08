package response

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGetHBInfo struct {
	response.ResponsePayment

	XMLName     xml.Name `xml:"xml"`
	Text        string   `xml:",chardata"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDes  string   `xml:"err_code_des"`
	MchBillno   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	DetailID    string   `xml:"detail_id"`
	Status      string   `xml:"status"`
	SendType    string   `xml:"send_type"`
	HbType      string   `xml:"hb_type"`
	TotalNum    string   `xml:"total_num"`
	TotalAmount string   `xml:"total_amount"`
	SendTime    string   `xml:"send_time"`
	Hblist      struct {
		Text   string `xml:",chardata"`
		Hbinfo struct {
			Text    string `xml:",chardata"`
			OpenID  string `xml:"openid"`
			Amount  string `xml:"amount"`
			RcvTime string `xml:"rcv_time"`
		} `xml:"hbinfo"`
	} `xml:"hblist"`
}
