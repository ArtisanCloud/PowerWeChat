package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMeetingCreate struct {
	*response.ResponseWork

	MeetingID int `json:"meetingid"`
}
