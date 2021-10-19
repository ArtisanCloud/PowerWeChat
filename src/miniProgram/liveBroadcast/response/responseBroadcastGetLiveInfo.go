package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseBroadcastGetLiveInfo struct {
	*response.ResponseMiniProgram

	Total int `json:"total"`
	RoomInfo []*power.HashMap `json:"room_info"`


}
