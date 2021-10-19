package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCustomerStrategyGet struct {
	*response.ResponseWork

	Strategy *power.HashMap `json:"momentStrategy"`
}
