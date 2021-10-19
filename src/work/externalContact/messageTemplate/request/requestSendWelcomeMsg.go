package request

import "github.com/ArtisanCloud/powerwechat/src/kernel/power"

type RequestSendWelcomeMsg struct {
	WelcomeCode string           `json:"welcome_code"`
	Text        *TextOfMessage   `json:"text"`
	Attachments []*power.HashMap `json:"attachments"`
}
