package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMeetingRoomAdd struct {
	*response.ResponseWork

	MeetingRoomID int `json:"meetingroom_id"`
}
