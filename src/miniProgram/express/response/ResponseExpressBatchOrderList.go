package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseExpressBatchOrderList struct {
	*response.ResponseMiniProgram

	OrderList []*power.HashMap `json:"order_list"`
}
