package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMomentStrategyGetRange struct {
	*response.ResponseWork

	Range      []*power.HashMap `json:"range"`
	NextCursor string           `json:"next_cursor"`
}
