package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseBroadcastGetLiveInfo struct {
	*response.ResponseMiniProgram

	Total int `json:"total"`
	RoomInfo []*power.HashMap `json:"room_info"`


}
