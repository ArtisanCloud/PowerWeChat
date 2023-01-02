package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyList struct {
	response.ResponseWork

	Strategy []*power.HashMap `json:"momentStrategy"`
}
