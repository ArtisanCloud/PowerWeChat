package request

type RequestAccountServiceSendMsgOnEvent struct {
	Code    string                       `json:"code"`
	MsgID   string                       `json:"msgid"`
	MsgType string                       `json:"msgtype"`
	Text    RequestAccountServiceMsgText `json:"text,omitempty"`
	Menu    RequestAccountServiceMsgMenu `json:"msgmenu,omitempty"`
}
