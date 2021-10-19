package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseBroadcastGoodsAudit struct {
	*response.ResponseMiniProgram

	AuditID int `json:"auditId"`
}
