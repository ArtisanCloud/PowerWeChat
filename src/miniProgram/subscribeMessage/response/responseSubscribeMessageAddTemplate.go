package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	*response.ResponseMiniProgram
	PriTmplID   string `json:"priTmplId"`
}
