package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	*response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
