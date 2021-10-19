package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMeetingRoomAdd struct {
	*response.ResponseWork

	MeetingRoomID int `json:"meetingroom_id"`
}
