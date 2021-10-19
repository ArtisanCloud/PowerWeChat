package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMeetingGetUserMeetingID struct {
	*response.ResponseWork

	NextCursor    string   `json:"next_cursor"`
	MeetingidList []string `json:"meetingid_list"`
}
