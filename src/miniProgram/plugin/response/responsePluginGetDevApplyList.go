package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponsePluginGetDevApplyList struct {
	*response.ResponseMiniProgram
	ApplyList []*power.HashMap `json:"apply_list"`
}
