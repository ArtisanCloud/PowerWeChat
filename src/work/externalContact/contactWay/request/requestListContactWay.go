package request

type RequestListContactWay struct {
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Cursor    string `json:"cursor"`
	Limit     int    `json:"limit"`
}
