package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	*response.ResponseMiniProgram
	PriTmplID   string `json:"priTmplId"`
}
