package request

type RequestMeetingRoomBook struct {
	MeetingRoomID int      `json:"meetingroom_id"`
	Subject       string   `json:"subject"`
	StartTime     int64    `json:"start_time"`
	EndTime       int64    `json:"end_time"`
	Booker        string   `json:"booker"`
	Attendees     []string `json:"attendees"`
}
