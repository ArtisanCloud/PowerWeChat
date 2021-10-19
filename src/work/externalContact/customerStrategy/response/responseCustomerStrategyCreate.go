package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	*response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
