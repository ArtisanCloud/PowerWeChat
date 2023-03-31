package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseSecurityMsgCheckASync struct {
	response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
	Detail  []struct {
		Strategy string `json:"momentStrategy"`
		ErrCode  int    `json:"errcode"`
		Suggest  string `json:"suggest"`
		Label    int    `json:"label"`
		Prob     int    `json:"prob,omitempty"`
		Level    int    `json:"level,omitempty"`
		Keyword  string `json:"keyword,omitempty"`
	} `json:"detail"`
	Result struct {
		Suggest string `json:"suggest"`
		Label   int    `json:"label"`
	} `json:"result"`
}
