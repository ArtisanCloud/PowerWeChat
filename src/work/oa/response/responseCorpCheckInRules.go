package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCorpCheckInRules struct {
	*response.ResponseWork
	Group []*power.HashMap `json:"group"`
}
