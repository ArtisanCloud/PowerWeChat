package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseUnitfy struct {
	*response.ResponsePayment

	PrepayID string `json:"prepay_id"`
}
