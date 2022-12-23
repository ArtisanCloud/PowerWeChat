package request

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
)

type RequestGroupWelcomeTemplateAdd struct {
	Text        *request.TextOfMessage `json:"text"`
	Image       *request.Image         `json:"image"`
	Link        *request.Link          `json:"link"`
	MiniProgram *request.MiniProgram   `json:"miniprogram"`
	File        *request.File          `json:"file"`
	Video       *request.Video         `json:"video"`
	AgentID     int64                  `json:"agentid"`
	Notify      int                    `json:"notify"`
}
