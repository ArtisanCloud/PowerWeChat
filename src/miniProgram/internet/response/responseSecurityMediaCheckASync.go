package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseSecurityMediaCheckASync struct {
	*response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
}
