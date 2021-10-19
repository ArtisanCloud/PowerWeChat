package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	*response.ResponseWork
	IDs []string `json:"ids"`
}


