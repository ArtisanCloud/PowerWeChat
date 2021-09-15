package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseExpressCancelOrder struct {
	*response.ResponseMiniProgram

	DeliveryResultCode int    `json:"delivery_resultcode"` //  0,
	DeliveryResultMsg  string `json:"delivery_resultmsg"`  //  ""

}
