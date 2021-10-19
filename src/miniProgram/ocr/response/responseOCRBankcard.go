package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseOCRBankcard struct {
	*response.ResponseMiniProgram
	Number string `json:"number"`
}
