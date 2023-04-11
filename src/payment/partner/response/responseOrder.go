package response

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"time"
)

// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_2.shtml

type Amount struct {
	Currency      string `json:"currency"`
	PayerCurrency string `json:"payer_currency"`
	PayerTotal    int    `json:"payer_total"`
	Total         int    `json:"total"`
}

type Payer struct {
	SpOpenid  string `json:"sp_openid"`
	SubOpenid string `json:"sub_openid"`
}

type ResponseOrder struct {
	response.ResponsePayment

	Amount          *Amount           `json:"amount"`
	Attach          string            `json:"attach"`
	BankType        string            `json:"bank_type"`
	OutTradeNo      string            `json:"out_trade_no"`
	Payer           *Payer            `json:"payer"`
	PromotionDetail []*object.HashMap `json:"promotion_detail"`
	SpAppId         string            `json:"sp_appid"`
	SpMchId         string            `json:"sp_mchid"`
	SubAppId        string            `json:"sub_appid"`
	SubMchId        string            `json:"sub_mchid"`
	SuccessTime     time.Time         `json:"success_time"`
	TradeState      string            `json:"trade_state"`
	TradeStateDesc  string            `json:"trade_state_desc"`
	TradeType       string            `json:"trade_type"`
	TransactionId   string            `json:"transaction_id"`
}
