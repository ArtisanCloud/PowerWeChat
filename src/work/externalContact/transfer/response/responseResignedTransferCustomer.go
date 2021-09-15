package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseResignedTransferCustomer struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}
