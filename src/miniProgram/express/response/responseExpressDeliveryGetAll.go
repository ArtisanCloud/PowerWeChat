package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseExpressDeliveryGetAll struct {
	*response.ResponseMiniProgram

	Count int               `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
