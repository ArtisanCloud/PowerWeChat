package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml

type ResponseReserve struct {
	*response.ResponsePayment
}
