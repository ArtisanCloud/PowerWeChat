package request

type RequestDialGetDialRecord struct {
	MeetingID       int64 `json:"start_time"`
	Title           int64 `json:"end_time"`
	MeetingStart    int   `json:"offset"`
	MeetingDuration int   `json:"limit"`
}
