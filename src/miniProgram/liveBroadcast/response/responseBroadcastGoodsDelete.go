package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseBroadcastGoodsDelete struct {
	*response.ResponseMiniProgram

	Total int              `json:"total"`
	Goods []*power.HashMap `json:"goods"`
}
