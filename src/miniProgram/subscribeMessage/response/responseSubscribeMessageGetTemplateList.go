package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseSubscribeMessageGetTemplateList struct {
	*response.ResponseMiniProgram
	Data   []*power.HashMap `json:"data"`
}
