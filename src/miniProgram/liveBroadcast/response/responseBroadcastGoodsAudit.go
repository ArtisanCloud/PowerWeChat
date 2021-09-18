package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGoodsAudit struct {
	*response.ResponseMiniProgram

	AuditID int `json:"auditId"`
}
