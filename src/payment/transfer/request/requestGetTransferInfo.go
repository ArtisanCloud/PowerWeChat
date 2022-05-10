package request

import "encoding/xml"

type RequestTransferToBalance struct {
	XMLName          xml.Name `xml:"xml"`
	Text             string   `xml:",chardata"`
	MchAppID         string   `xml:"mch_appid" json:"mch_appid"`
	MchID            string   `xml:"mchid" json:"mchid"`
	NonceStr         string   `xml:"nonce_str" json:"nonce_str"`
	PartnerTradeNo   string   `xml:"partner_trade_no" json:"partner_trade_no"`
	Openid           string   `xml:"openid" json:"openid"`
	CheckName        string   `xml:"check_name" json:"check_name"`
	ReUserName       string   `xml:"re_user_name" json:"re_user_name"`
	Amount           int64    `xml:"amount" json:"amount"`
	Desc             string   `xml:"desc" json:"desc"`
	SpBillCreateIP   string   `xml:"spbill_create_ip" json:"spbill_create_ip"`
	Scene            string   `xml:"scene" json:"scene"`
	BrandID          string   `xml:"brand_id" json:"brand_id"`
	FinderTemplateID string   `xml:"finder_template_id" json:"finder_template_id"`
}
