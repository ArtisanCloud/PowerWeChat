package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressGetOrder struct {
	response.ResponseMiniProgram

	PrintHTML   string             `json:"print_html"`   //  "客户密码不正确"
	WaybillData []*power.StringMap `json:"waybill_data"` //
	DeliveryID  string             `json:"delivery_id"`  //  "",
	WaybillID   string             `json:"waybill_id"`   //   "123456789",
	OrderID     string             `json:"order_id"`     //   "01234567890123456789",
	OrderStatus int                `json:"order_status"` //  0

}
