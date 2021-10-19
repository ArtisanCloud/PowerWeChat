package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMeetingRoomList struct {
	*response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}
