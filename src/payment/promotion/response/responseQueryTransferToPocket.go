package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// https://work.weixin.qq.com/api/doc/90000/90135/90278

type ResponseQueryTransferToPocket struct {
	response.ResponsePayment

	MchID          string `json:"mch_id"`
	AppID          string `json:"appid"`
	DetailID       string `json:"detail_id"`
	PartnerTradeNO string `json:"partner_trade_no"`
	Status         string `json:"status"`
	PaymentAmount  int    `json:"payment_amount"`
	OpenID         string `json:"openid"`
	TransferTime   string `json:"transfer_time"`
	TransferName   string `json:"transfer_name"`
	Desc           string `json:"desc"`
}
