package response

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseImmediateDeliveryDelivery struct {
	*response.ResponseMiniProgram

	List []*object.HashMap `json:"list"`
}
