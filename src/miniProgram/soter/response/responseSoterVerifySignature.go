package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	*response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
