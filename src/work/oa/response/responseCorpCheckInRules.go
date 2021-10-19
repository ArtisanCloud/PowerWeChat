package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCorpCheckInRules struct {
	*response.ResponseWork
	Group []*power.HashMap `json:"group"`
}
