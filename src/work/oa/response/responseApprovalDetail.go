package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseApprovalDetail struct {
	*response.ResponseWork
	Info power.HashMap `json:"info"`
}
