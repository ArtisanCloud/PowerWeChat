package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseBroadcastGoodsAdd struct {
	*response.ResponseMiniProgram

	GoodsID int64 `json:"goodsId"`
	AuditID int64 `json:"auditId"`


}
