package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	*response.ResponseWork
	IDs []string `json:"ids"`
}


