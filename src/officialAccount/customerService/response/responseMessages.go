package response

type Record struct {
	OpenID   string `json:"openid"`
	OperCode int    `json:"opercode"`
	Text     string `json:"text"`
	Time     int    `json:"time"`
	Worker   string `json:"worker"`
}

type ResponseMessages struct {
	RecordList []*Record `json:"recordlist"`
	Number int `json:"number"`
	MsgID  int `json:"msgid"`
}
