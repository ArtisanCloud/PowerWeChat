package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseServiceMarketInvoceService struct {
	response.ResponseMiniProgram

	Data string `json:"data"`
}
