package request

type RequestSendWelcomeMsg struct {
	WelcomeCode string                     `json:"welcome_code"`
	Text        *TextOfMessage             `json:"text"`
	Attachments []MessageTemplateInterface `json:"attachments"`
}
