package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCustomerStrategyGetRange struct {
	*response.ResponseWork

	Range []*power.HashMap `json:"range"`
}
