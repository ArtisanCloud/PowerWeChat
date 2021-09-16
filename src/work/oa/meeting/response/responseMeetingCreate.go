package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMeetingCreate struct {
	*response.ResponseWork

	MeetingID int `json:"meetingid"`
}
