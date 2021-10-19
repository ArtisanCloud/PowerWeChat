package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCustomerStrategyList struct {
	*response.ResponseWork

	Strategy []*power.HashMap `json:"momentStrategy"`
}
