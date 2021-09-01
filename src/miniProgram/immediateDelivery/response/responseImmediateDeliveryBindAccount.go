package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	response.ResponseMiniProgram
	ShopList []*object.HashMap `json:"shop_list"`
}
