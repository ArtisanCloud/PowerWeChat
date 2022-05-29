package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseOCRBankcard struct {
	*response.ResponseMiniProgram
	Number string `json:"number"`
}
