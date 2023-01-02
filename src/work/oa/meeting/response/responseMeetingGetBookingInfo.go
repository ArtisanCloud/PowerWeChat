package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingGetBookingInfo struct {
	response.ResponseWork

	CreatorUserID          string         `json:"creator_userid"`
	Title                  string         `json:"title"`
	ReserveMeetingStart    int64          `json:"reserve_meeting_start"`
	ReserveMeetingDuration int64          `json:"reserve_meeting_duration"`
	MeetingStart           int64          `json:"meeting_start"`
	MeetingDuration        int64          `json:"meeting_duration"`
	Description            string         `json:"description"`
	MainDepartment         int            `json:"main_department"`
	Type                   int            `json:"type"`
	Status                 int            `json:"status"`
	RemindTime             int64          `json:"remind_time"`
	Attendees              *power.HashMap `json:"attendees"`
}
