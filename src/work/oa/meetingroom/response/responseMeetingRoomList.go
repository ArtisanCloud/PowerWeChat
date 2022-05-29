package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMeetingRoomList struct {
	*response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}
