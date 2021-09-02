package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOCRBankcard struct {
	*response.ResponseMiniProgram
	Number string `json:"number"`
}
