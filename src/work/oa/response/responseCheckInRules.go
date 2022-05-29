package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCheckInRules struct {
	*response.ResponseWork
	Info []*power.HashMap `json:"info"`
}
