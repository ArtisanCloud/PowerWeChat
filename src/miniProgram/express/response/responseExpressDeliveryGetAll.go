package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseExpressDeliveryGetAll struct {
	*response.ResponseMiniProgram

	Count int               `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
