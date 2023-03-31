package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressDeliveryGetAll struct {
	response.ResponseMiniProgram

	Count int              `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
