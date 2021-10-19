package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseNearbyPoiGetList struct {
	*response.ResponseMiniProgram
	Data *power.HashMap `json:"data"`
}
