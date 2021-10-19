package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseResignedTransferResult struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}
