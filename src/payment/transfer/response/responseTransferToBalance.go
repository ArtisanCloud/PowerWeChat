package response

type ResponseTransferToBalance struct {
	*ResponseTransfer

	MchAppid       string   `xml:"mch_appid"`
	Mchid          string   `xml:"mchid"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	PaymentNo      string   `xml:"payment_no"`
	PaymentTime    string   `xml:"payment_time"`
}
