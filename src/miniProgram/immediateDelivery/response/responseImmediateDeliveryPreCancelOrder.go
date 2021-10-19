package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseImmediateDeliveryPreCancelOrder struct {
	*response.ResponseMiniProgram

	DeductFee int    `json:"deduct_fee"` // 5,
	Desc      string `json:"desc"`       // "blablabla",

}
