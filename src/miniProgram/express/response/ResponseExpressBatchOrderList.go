package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
)

type ResponseExpressBatchOrderList struct {
	OrderList []*power.HashMap `json:"order_list"`
}
