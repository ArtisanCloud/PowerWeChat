package request

type RequestGetGroupMsgListV2 struct {
	ChatType   string  `json:"chat_type"`
	StartTime  int64   `json:"start_time"`
	EndTime    int64   `json:"end_time"`
	Creator    *string `json:"creator"`
	FilterType *int    `json:"filter_type"`
	Limit      int     `json:"limit"`
	Cursor     string  `json:"cursor"`
}
