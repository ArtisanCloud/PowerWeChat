package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponsePluginGetDevApplyList struct {
	*response.ResponseMiniProgram
	ApplyList []*power.HashMap `json:"apply_list"`
}
