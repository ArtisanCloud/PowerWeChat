package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMomentStrategyList struct {
	*response.ResponseWork

	StrategyID int64 `json:"strategy_id"`
}
