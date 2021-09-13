package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	*response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
