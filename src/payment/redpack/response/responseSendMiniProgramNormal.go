package response

import "encoding/xml"

type ResponseSendMiniProgramNormal struct {
	XMLName     xml.Name `xml:"xml"`
	Text        string   `xml:",chardata"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDes  string   `xml:"err_code_des"`
	MchBillNO   string   `xml:"mch_billno"`
	MchID       string   `xml:"mch_id"`
	WXAppID     string   `xml:"wxappid"`
	ReOpenID    string   `xml:"re_openid"`
	TotalAmount string   `xml:"total_amount"`
	SendListID  string   `xml:"send_listid"`
	Package     string   `xml:"package"`
}

