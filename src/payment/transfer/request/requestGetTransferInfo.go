package request

type RequestTransferToBalance struct {
	MchAppID         string `xml:"mch_appid" json:"mch_appid"`
	MchID            string `xml:"mchid" json:"mchid"`
	DeviceInfo       string `xml:"device_info" json:"device_info"`
	NonceStr         string `xml:"nonce_str" json:"nonce_str"`
	PartnerTradeNo   string `xml:"partner_trade_no" json:"partner_trade_no"`
	OpenID           string `xml:"openid" json:"openid"`
	CheckName        string `xml:"check_name" json:"check_name"`
	ReUserName       string `xml:"re_user_name" json:"re_user_name"`
	Amount           int    `xml:"amount" json:"amount"`
	Desc             string `xml:"desc" json:"desc"`
	SpBillCreateIP   string `xml:"spbill_create_ip" json:"spbill_create_ip"`
	Scene            string `xml:"scene" json:"scene"`
	BrandID          int    `xml:"brand_id" json:"brand_id"`
	FinderTemplateID string `xml:"finder_template_id" json:"finder_template_id"`
}
