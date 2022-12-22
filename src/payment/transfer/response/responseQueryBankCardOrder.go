package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetTransferInfo struct {
	response.ResponsePayment

	MchID          string `xml:"mch_id"`
	AppID          string `xml:"appid"`
	DetailID       string `xml:"detail_id"`
	PartnerTradeNO string `xml:"partner_trade_no"`
	Status         string `xml:"status"`
	PaymentAmount  string `xml:"payment_amount"`
	OpenID         string `xml:"openid"`
	TransferTime   string `xml:"transfer_time"`
	TransferName   string `xml:"transfer_name"`
	Desc           string `xml:"desc"`
}
