package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseBroadcastGoodsAdd struct {
	*response.ResponseMiniProgram

	GoodsID string `json:"goodsId"`
	AuditID int64 `json:"auditId"`

}
