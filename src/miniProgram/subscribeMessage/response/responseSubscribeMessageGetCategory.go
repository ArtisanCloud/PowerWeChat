package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseSubscribeMessageGetCategory struct {
	*response.ResponseMiniProgram
	Data   []*power.HashMap `json:"data"`
}
