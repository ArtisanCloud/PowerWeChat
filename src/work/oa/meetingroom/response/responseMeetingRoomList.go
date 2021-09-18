package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMeetingRoomList struct {
	*response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}
