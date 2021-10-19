package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	*response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
