package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMeetingRoomAdd struct {
	*response.ResponseWork

	MeetingRoomID int `json:"meetingroom_id"`
}
