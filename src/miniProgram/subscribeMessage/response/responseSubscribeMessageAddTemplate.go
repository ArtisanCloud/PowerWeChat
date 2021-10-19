package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	*response.ResponseMiniProgram
	PriTmplID   string `json:"priTmplId"`
}
