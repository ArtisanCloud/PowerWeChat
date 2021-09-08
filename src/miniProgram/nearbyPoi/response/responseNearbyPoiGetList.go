package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseNearbyPoiGetList struct {
	*response.ResponseMiniProgram
	Data *power.HashMap `json:"data"`
}
