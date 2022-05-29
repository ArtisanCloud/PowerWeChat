package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMeetingRoomGetBook struct {
	*response.ResponseWork

	MeetingID  int `json:"meeting_id"`
	ScheduleID int `json:"schedule_id"`
}
