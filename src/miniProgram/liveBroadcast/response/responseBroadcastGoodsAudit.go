package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseBroadcastGoodsAudit struct {
	*response.ResponseMiniProgram

	AuditID int `json:"auditId"`
}
