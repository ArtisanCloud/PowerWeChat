package request

import "github.com/ArtisanCloud/power-wechat/src/work/externalContact/messageTemplate/request"

type RequestGroupWelcomeTemplateEdit struct {
	TemplateID  string                 `json:"template_id"`
	Text        *request.TextOfMessage `json:"text"`
	Image       *request.Image         `json:"image"`
	Link        *request.Link          `json:"link"`
	MiniProgram *request.MiniProgram   `json:"miniprogram"`
	File        *request.File          `json:"file"`
	Video       *request.Video         `json:"video"`
	AgentID     int64                  `json:"agentid"`
	Notify      int                    `json:"notify"`
}
