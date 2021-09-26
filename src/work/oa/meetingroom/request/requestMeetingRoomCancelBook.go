package request

type RequestMeetingRoomCancelBook struct {
	MeetingID    string `json:"meeting_id"`
	KeepSchedule int    `json:"keep_schedule"`
}
