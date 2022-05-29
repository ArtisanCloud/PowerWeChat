package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCustomerStrategyGetRange struct {
	*response.ResponseWork

	Range []*power.HashMap `json:"range"`
}
