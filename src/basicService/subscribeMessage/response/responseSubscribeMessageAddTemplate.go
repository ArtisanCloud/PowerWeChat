package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	response.ResponseMiniProgram
	PriTmplID string `json:"priTmplId"`
}
