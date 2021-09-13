package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseServiceMarketInvoceService struct {
	*response.ResponseMiniProgram

	Data string `json:"data"`
}
