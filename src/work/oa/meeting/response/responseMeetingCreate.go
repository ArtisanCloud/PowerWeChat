package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseMeetingCreate struct {
	*response.ResponseWork

	MeetingID int `json:"meetingid"`
}
