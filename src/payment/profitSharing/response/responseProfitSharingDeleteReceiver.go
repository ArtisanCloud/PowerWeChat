package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml

type ResponseProfitSharingDeleteReceiver struct {
	response.ResponsePayment

	Type    string `json:"type"`
	Account string `json:"accountService"`
}
