package response

import (
	"encoding/xml"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseReserve struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	*response.ResponsePayment
	AppID    string `xml:"appid"`
	MchID    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`
}
