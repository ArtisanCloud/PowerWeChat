package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	*response.ResponseMiniProgram

	ShopList []*power.HashMap `json:"shop_list"`
}
