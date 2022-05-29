package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	*response.ResponseMiniProgram

	ShopList []*power.HashMap `json:"shop_list"`
}
