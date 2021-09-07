package request

import "github.com/ArtisanCloud/power-wechat/src/kernel/power"

type WeAppTemplateMsg struct {
	TemplateID      string         `json:"template_id"`
	Page            string         `json:"page"`
	FormID          string         `json:"form_id"`
	Data            *power.HashMap `json:"data"`
	EmphasisKeyword string         `json:"emphasis_keyword"`
}

type MPTemplateMsg struct {
	AppID       string         `json:"appid"`
	TemplateID  string         `json:"template_id"`
	URL         string         `json:"url"`
	MiniProgram string         `json:"miniprogram"`
	Data        *power.HashMap `json:"data"`
}
