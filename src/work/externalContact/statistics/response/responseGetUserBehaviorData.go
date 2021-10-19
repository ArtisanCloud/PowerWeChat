package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGetUserBehaviorData struct {
	*response.ResponseWork

	MomentList []*power.HashMap `json:"behavior_data"`
}
