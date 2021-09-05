package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseSubscribeMessageGetPubTemplateTitleList struct {
	*response.ResponseMiniProgram
	Count   string `json:"count"`
	Data   []*power.HashMap `json:"data"`
}
