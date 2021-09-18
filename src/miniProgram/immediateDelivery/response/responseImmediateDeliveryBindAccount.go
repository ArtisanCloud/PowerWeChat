package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	*response.ResponseMiniProgram

	ShopList []*power.HashMap `json:"shop_list"`
}
