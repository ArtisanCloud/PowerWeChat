package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	*response.ResponseWork
	IDs []string `json:"ids"`
}
