package response

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml

type ResponseProfitSharingGetBills struct {
	DownloadURL string `json:"download_url"`
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
}
