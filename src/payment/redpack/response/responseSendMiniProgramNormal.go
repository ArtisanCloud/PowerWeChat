package response

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseSendMiniProgramNormal struct {
	response.ResponsePayment

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
	TotalAmount int      `xml:"total_amount"`
	SendListID  string   `xml:"send_listid"`
	Package     string   `xml:"package"`
}
