package response

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_9.shtml

type ResponseProfitSharingDeleteReceiver struct {
	Type    string `json:"type"`
	Account string `json:"account"`
}
