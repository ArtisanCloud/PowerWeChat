package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMomentStrategyCreate struct {
	*response.ResponseWork

	Strategy   []*power.HashMap `json:"strategy"`
	NextCursor string           `json:"next_cursor"`
}
