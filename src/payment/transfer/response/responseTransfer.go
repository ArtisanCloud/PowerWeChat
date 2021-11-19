package response

import (
	"encoding/xml"
)

// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3

type ResponseTransfer struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	ResultCode string   `xml:"result_code"`
}
