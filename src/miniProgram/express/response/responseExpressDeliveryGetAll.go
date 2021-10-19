package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseExpressDeliveryGetAll struct {
	*response.ResponseMiniProgram

	Count int               `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
