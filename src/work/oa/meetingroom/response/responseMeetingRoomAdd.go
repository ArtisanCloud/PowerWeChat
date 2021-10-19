package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMeetingRoomAdd struct {
	*response.ResponseWork

	MeetingRoomID int `json:"meetingroom_id"`
}
