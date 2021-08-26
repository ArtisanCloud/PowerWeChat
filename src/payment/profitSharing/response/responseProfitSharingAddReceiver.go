package response

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml

type ResponseProfitSharingAddReceiver struct {
	Type         string `json:"type"`
	Account      string `json:"account"`
	Name         string `json:"name"`
	RelationType string `json:"relation_type"`
}
