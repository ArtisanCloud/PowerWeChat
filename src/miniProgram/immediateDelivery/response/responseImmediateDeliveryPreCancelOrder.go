package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseImmediateDeliveryPreCancelOrder struct {
	*response.ResponseMiniProgram

	DeductFee int    `json:"deduct_fee"` // 5,
	Desc      string `json:"desc"`       // "blablabla",

}
