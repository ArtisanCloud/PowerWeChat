package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCorpCheckInRules struct {
	*response.ResponseWork
	Group []*power.HashMap `json:"group"`
}
