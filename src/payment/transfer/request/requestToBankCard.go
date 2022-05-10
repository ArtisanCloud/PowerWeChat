package request

import "encoding/xml"

type RequestToBankCard struct {
	XMLName        xml.Name `xml:"xml"`
	Text           string   `xml:",chardata"`
	Amount         int64   `xml:"amount" json:"amount"`
	BankCode       string   `xml:"bank_code" json:"bank_code"`
	Desc           string   `xml:"desc" json:"desc"`
	EncBankNO      string   `xml:"enc_bank_no" json:"enc_bank_no"`
	EncTrueName    string   `xml:"enc_true_name" json:"enc_true_name"`
	MchID          int64   `xml:"mch_id" json:"mch_id"`
	NonceStr       string   `xml:"nonce_str" json:"nonce_str"`
	PartnerTradeNO int64    `xml:"partner_trade_no" json:"partner_trade_no"`
}