package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3

type ResponseTransfer struct {
	*response.ResponsePayment
}
