package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMeetingRoomGetBook struct {
	*response.ResponseWork

	MeetingID int `json:"meeting_id"`
	ScheduleID int `json:"schedule_id"`

}
