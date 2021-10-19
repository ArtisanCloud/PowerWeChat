package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseMeetingCreate struct {
	*response.ResponseWork

	MeetingID int `json:"meetingid"`
}
