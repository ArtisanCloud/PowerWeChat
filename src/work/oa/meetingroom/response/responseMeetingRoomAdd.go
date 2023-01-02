package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomAdd struct {
	response.ResponseWork

	MeetingRoomID int `json:"meetingroom_id"`
}
