package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	*response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
