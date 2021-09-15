package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetUserBehaviorData struct {
	*response.ResponseWork

	MomentList []*power.HashMap `json:"behavior_data"`
}
