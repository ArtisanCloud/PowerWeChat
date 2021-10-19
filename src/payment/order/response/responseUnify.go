package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseUnitfy struct {
	*response.ResponsePayment

	PrepayID string `json:"prepay_id"`
}
