package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMomentStrategyGetRange struct {
	*response.ResponseWork

	Range      []*power.HashMap `json:"range"`
	NextCursor string           `json:"next_cursor"`
}
