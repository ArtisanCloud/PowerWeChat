package request

type RequestMeetingRoomGetBookingInfo struct {
	MeetingroomID int    `json:"meetingroom_id"`
	StartTime     int    `json:"start_time"`
	EndTime       int    `json:"end_time"`
	City          string `json:"city"`
	Building      string `json:"building"`
	Floor         string `json:"floor"`
}
