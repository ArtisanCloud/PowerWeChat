package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseServiceMarketInvoceService struct {
	*response.ResponseMiniProgram

	Data string `json:"data"`
}
