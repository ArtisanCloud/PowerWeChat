package request

type Attachment struct {
	MsgType     string       `json:"msgtype"`
	Image       *Image       `json:"image,omitempty"`
	Link        *Link        `json:"link,omitempty"`
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"`
	Video       *Video       `json:"video,omitempty"`
	File        *File        `json:"file,omitempty"`
}

type RequestSendWelcomeMsg struct {
	WelcomeCode string         `json:"welcome_code"`
	Text        *TextOfMessage `json:"text"`
	Attachments []*Attachment  `json:"attachments"`
}
