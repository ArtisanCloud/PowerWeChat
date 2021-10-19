package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMomentStrategyCreate struct {
	*response.ResponseWork

	Strategy   []*power.HashMap `json:"strategy"`
	NextCursor string           `json:"next_cursor"`
}