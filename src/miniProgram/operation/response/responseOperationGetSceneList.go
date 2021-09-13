package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOperationGetSceneList struct {
	*response.ResponseMiniProgram
	Scene []*power.HashMap `json:"scene"`
}
