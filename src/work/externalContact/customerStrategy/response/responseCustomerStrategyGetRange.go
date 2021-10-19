package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCustomerStrategyGetRange struct {
	*response.ResponseWork

	Range []*power.HashMap `json:"range"`
}
