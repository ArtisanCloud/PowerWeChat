package response

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseImmediateDeliveryDelivery struct {
	*response.ResponseMiniProgram

	List []*object.HashMap `json:"list"`
}
