package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseServiceMarketInvoceService struct {
	*response.ResponseMiniProgram

	Data string `json:"data"`
}
