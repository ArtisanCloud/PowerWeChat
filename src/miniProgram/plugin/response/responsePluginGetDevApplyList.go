package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponsePluginGetDevApplyList struct {
	*response.ResponseMiniProgram
	ApplyList []*power.HashMap `json:"apply_list"`
}
