package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGoodsAudit struct {
	response.ResponseMiniProgram

	AuditID int `json:"auditId"`
}
