package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	*response.ResponseMiniProgram

	IsOK string `json:"is_ok"`
}
