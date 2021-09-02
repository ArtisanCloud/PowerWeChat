package response

import (
	"github.com/ArtisanCloud/go-libs/object"
)

type ResponseExpressGetPath struct {
	OpenID       string              `json:"openid"`        //   "openid",
	DeliveryID   string              `json:"delivery_id"`   //  "",
	WaybillID    string              `json:"waybill_id"`    //   "123456789",
	PathItemNum  int                 `json:"path_item_num"` //  3,
	PathItemList []*object.StringMap `json:"path_item_list"`
}
