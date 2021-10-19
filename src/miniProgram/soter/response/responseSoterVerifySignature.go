package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	*response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
