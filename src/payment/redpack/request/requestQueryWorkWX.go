package request

type RequestQueryWorkWX struct {

	NonceStr  string   `xml:"nonce_str"`
	Sign      string   `xml:"sign"`
	MchBillNO string   `xml:"mch_billno"`
	MchID     string   `xml:"mch_id"`
	Appid     string   `xml:"appid"`
}
