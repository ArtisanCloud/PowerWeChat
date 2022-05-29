package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseSubscribeMessageGetPubTemplateTitleList struct {
	*response.ResponseMiniProgram
	Count int              `json:"count"`
	Data  []*power.HashMap `json:"data"`
}
