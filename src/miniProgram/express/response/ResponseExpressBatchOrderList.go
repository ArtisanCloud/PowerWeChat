package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
)

type ResponseExpressBatchOrderList struct {
	OrderList []*power.HashMap `json:"order_list"`
}
