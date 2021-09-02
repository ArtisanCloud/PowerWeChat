package response

import (
	"github.com/ArtisanCloud/go-libs/object"
)

type ResponseExpressBatchOrderList struct {
	OrderList []*object.HashMap `json:"order_list"`
}
