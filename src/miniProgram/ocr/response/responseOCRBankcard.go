package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOCRBankcard struct {
	*response.ResponseMiniProgram
	Number string `json:"number"`
}
