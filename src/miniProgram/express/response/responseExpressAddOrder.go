package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseExpressAddOrder struct {
	response.ResponseMiniProgram

	DeliveryResultCode string              `json:"delivery_resultcode"` //  10002,
	DeliveryResultMsg  float64             `json:"delivery_resultmsg"`  //  "客户密码不正确"
	OrderID            int                 `json:"order_id"`            //   "01234567890123456789",
	WaybillID          string              `json:"waybill_id"`          //   "123456789",
	WaybillData        []*object.StringMap `json:"waybill_data"`        //

}
