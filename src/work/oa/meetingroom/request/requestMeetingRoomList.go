package request

type RequestMeetingRoomList struct {
	City      string `json:"city"`
	Building  string `json:"building"`
	Floor     string `json:"floor"`
	Equipment []int  `json:"equipment"`
}
