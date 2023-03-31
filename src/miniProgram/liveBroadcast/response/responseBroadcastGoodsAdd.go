package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseBroadcastGoodsAdd struct {
	response.ResponseMiniProgram

	GoodsID string `json:"goodsId"`
	AuditID int64  `json:"auditId"`
}
