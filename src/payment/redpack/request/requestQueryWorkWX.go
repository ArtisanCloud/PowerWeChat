package request

import "encoding/xml"

type RequestQueryWorkWX struct {
	XMLName   xml.Name `xml:"xml"`
	Text      string   `xml:",chardata"`
	NonceStr  string   `xml:"nonce_str"`
	Sign      string   `xml:"sign"`
	MchBillNO string   `xml:"mch_billno"`
	MchID     string   `xml:"mch_id"`
	Appid     string   `xml:"appid"`
}
