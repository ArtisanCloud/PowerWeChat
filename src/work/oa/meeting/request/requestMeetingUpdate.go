package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestMeetingUpdate struct {
	MeetingID       string         `json:"meetingid"`
	Title           string         `json:"title"`
	MeetingStart    int            `json:"meeting_start"`
	MeetingDuration int            `json:"meeting_duration"`
	Description     string         `json:"description"`
	Type            int            `json:"type"`
	RemindTime      int            `json:"remind_time"`
	Attendees       *power.HashMap `json:"attendees"`
}
