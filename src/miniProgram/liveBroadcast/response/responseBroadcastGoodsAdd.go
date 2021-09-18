package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseBroadcastGoodsAdd struct {
	*response.ResponseMiniProgram

	GoodsID string `json:"goodsId"`
	AuditID int64 `json:"auditId"`


}
