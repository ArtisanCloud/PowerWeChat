package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	*response.ResponseMiniProgram

	ShopList []*power.HashMap `json:"shop_list"`
}
