package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseSecurityMediaCheckASync struct {
	*response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
}
