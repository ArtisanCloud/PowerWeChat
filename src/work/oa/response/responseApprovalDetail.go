package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseApprovalDetail struct {
	*response.ResponseWork
	Info power.HashMap `json:"info"`
}
