package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMeetingRoomGetBook struct {
	*response.ResponseWork

	MeetingID int `json:"meeting_id"`
	ScheduleID int `json:"schedule_id"`

}
