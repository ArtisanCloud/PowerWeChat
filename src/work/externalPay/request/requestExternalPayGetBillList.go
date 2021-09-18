package request

type RequestExternalPayGetBillList struct {
	BeginTime   int64  `json:"begin_time"`
	EndTime     int64  `json:"end_time"`
	PayeeUserID string `json:"payee_userid"`
	Cursor      string `json:"cursor"`
	Limit       int    `json:"limit"`
}
