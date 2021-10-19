package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseSecurityMediaCheckASync struct {
	*response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
}
