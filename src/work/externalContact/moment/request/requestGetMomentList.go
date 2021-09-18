package request

type RequestGetMomentList struct {
	StartTime  int64  `json:"start_time" `
	EndTime    int64  `json:"end_time"`
	Creator    string `json:"creator"`
	FilterType int    `json:"filter_type"`
	Cursor     string `json:"cursor"`
	Limit      int    `json:"limit"`

}
