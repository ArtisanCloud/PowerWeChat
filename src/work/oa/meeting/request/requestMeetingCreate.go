package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestMeetingCreate struct {
	CreatorUserID   string         `json:"creator_userid"`
	Title           string         `json:"title"`
	MeetingStart    int            `json:"meeting_start"`
	MeetingDuration int            `json:"meeting_duration"`
	Description     string         `json:"description"`
	Type            int            `json:"type"`
	RemindTime      int            `json:"remind_time"`
	AgentID         int            `json:"agentid"`
	Attendees       *power.HashMap `json:"attendees"`
}
