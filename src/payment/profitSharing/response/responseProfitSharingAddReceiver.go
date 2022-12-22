package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_1_8.shtml

type ResponseProfitSharingAddReceiver struct {
	response.ResponsePayment

	Type           string `json:"type"`
	Account        string `json:"account"`
	Name           string `json:"name,omitempty"`
	RelationType   string `json:"relation_type,omitempty"`
	CustomRelation string `json:"custom_relation,omitempty"`
}
