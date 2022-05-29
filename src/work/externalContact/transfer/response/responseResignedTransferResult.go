package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseResignedTransferResult struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}
