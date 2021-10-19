package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	*response.ResponseWork
	IDs []string `json:"ids"`
}


