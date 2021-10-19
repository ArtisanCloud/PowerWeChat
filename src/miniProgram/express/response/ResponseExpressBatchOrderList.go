package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
)

type ResponseExpressBatchOrderList struct {
	OrderList []*power.HashMap `json:"order_list"`
}
