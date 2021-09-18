package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCustomerStrategyGet struct {
	*response.ResponseWork

	Strategy *power.HashMap `json:"strategy"`
}
