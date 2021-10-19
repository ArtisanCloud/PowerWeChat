package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGetUserBehaviorData struct {
	*response.ResponseWork

	MomentList []*power.HashMap `json:"behavior_data"`
}
