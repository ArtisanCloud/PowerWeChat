package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseOperationGetSceneList struct {
	*response.ResponseMiniProgram
	Scene []*power.HashMap `json:"scene"`
}
