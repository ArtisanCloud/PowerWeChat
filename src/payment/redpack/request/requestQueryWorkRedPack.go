package request

import "encoding/xml"

type RequestQueryWorkRedPack struct {
	XMLName   xml.Name `xml:"xml"`
	Text      string   `xml:",chardata"`
	NonceStr  string   `xml:"nonce_str"`
	Sign      string   `xml:"sign"`
	MchBillno string   `xml:"mch_billno"`
	MchID     string   `xml:"mch_id"`
	Appid     string   `xml:"appid"`
}

