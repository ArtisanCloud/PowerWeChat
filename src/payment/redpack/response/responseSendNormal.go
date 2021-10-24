package response

import "encoding/xml"

type ResponseSendNormal struct {
	XMLName     xml.Name `xml:"xml"`
	Text        string   `xml:",chardata"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDes  string   `xml:"err_code_des"`
	MchBillno   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	Wxappid     string   `xml:"wxappid"`
	ReOpenid    string   `xml:"re_openid"`
	TotalAmount string   `xml:"total_amount"`
}

