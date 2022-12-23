package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseAuthCodeToOpenID struct {
	response.ResponsePayment

	AppID    string `json:"appid"`     // 是	String(32)	微信分配的公众账号ID
	MchID    string `json:"mch_id"`    // 是	String(32)	微信支付分配的商户号
	NonceStr string `json:"nonce_str"` // 是	String(32)	随机字符串，不长于32位
	Sign     string `json:"sign"`      // 是	String(32)	签名
	OpenID   string `json:"openid"`    // 是	String(128)	用户在商户appid下的唯一标识
}
