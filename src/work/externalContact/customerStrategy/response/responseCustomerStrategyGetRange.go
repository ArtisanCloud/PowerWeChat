package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCustomerStrategyGetRange struct {
	*response.ResponseWork

	Range []*power.HashMap `json:"range"`
}
