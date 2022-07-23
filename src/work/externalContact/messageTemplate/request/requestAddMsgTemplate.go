package request

type RequestAddMsgTemplate struct {
	ChatType       string                     `json:"chat_type"`
	ExternalUserID []string                   `json:"external_userid"`
	Sender         string                     `json:"sender"`
	Text           *TextOfMessage             `json:"text"`
	Attachments    []MessageTemplateInterface `json:"attachments"`
}
