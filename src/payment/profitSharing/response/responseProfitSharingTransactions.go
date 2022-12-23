package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_6.shtml

type ResponseProfitSharingTransaction struct {
	response.ResponsePayment

	TransactionID string `json:"transaction_id"`
	UnsplitAmount string `json:"unsplit_amount"`
}
