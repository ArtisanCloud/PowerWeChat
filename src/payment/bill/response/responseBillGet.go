package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_9.shtml

type ResponseBillGet struct {
	*response.ResponsePayment

	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadURL string `json:"download_url"`
}
