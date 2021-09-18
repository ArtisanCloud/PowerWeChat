package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGetLiveReplay struct {
	*response.ResponseMiniProgram

	Total int `json:"total"`
	LiveReplay []*power.HashMap `json:"live_replay"`


}
