package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOperationGetSceneList struct {
	*response.ResponseMiniProgram
	Scene []*power.HashMap `json:"scene"`
}
