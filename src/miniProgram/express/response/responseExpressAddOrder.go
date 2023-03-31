package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressAddOrder struct {
	response.ResponseMiniProgram

	DeliveryResultCode int                `json:"delivery_resultcode"` //  10002,
	DeliveryResultMsg  string             `json:"delivery_resultmsg"`  //  "客户密码不正确"
	OrderID            string             `json:"order_id"`            //   "01234567890123456789",
	WaybillID          string             `json:"waybill_id"`          //   "123456789",
	WaybillData        []*power.StringMap `json:"waybill_data"`        //

}
