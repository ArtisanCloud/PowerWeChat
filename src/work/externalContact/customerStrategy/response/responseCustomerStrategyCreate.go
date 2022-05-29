package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	*response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
