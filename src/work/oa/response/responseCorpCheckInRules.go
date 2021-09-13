package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpCheckInRules struct {
	*response.ResponseWork
	Group []*power.HashMap `json:"group"`
}
