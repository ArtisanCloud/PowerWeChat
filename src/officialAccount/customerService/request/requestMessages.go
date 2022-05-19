package request

type RequestMessages struct {
	StartTime int `json:"starttime"`
	EndTime   int `json:"endtime"`
	MsgID     int `json:"msgid"`
	Number    int `json:"number"`
}
