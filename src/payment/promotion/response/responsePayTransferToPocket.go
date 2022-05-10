package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

// https://work.weixin.qq.com/api/doc/90000/90135/90278

type ResponsePayTransferToPocket struct {
	*response.ResponsePayment

	AppID          string `json:"appid"`
	MchID          string `json:"mch_id"`
	DeviceInfo     string `json:"device_info"`
	NonceStr       string `json:"nonce_str"`
	PartnerTradeNO string `json:"partner_trade_no"`
	PaymentNO      string `json:"payment_no"`
	PaymentTime    string `json:"payment_time"`
}
