package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseImmediateDeliveryDelivery struct {
	*response.ResponseMiniProgram

	List []*object.HashMap `json:"list"`
}
