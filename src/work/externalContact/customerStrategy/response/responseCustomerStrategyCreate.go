package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	*response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
