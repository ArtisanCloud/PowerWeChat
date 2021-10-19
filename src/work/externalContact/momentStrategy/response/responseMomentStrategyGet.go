package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMomentStrategyGet struct {
	*response.ResponseWork

	Strategy   *power.HashMap `json:"strategy"`

}