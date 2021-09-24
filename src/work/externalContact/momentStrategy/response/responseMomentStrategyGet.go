package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMomentStrategyGet struct {
	*response.ResponseWork

	Strategy   *power.HashMap `json:"strategy"`

}