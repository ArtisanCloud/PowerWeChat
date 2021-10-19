package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseBroadcastGetLiveReplay struct {
	*response.ResponseMiniProgram

	Total int `json:"total"`
	LiveReplay []*power.HashMap `json:"live_replay"`


}
