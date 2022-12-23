package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestUniformMessageSend struct {
	ToUser           string            `json:"touser"`
	WeAppTemplateMsg *WeAppTemplateMsg `json:"weapp_template_msg,omitempty"`
	MpTemplateMsg    *MPTemplateMsg    `json:"mp_template_msg"`
}

type WeAppTemplateMsg struct {
	TemplateID      string         `json:"template_id"`
	Page            string         `json:"page"`
	FormID          string         `json:"form_id"`
	Data            *power.HashMap `json:"data"`
	EmphasisKeyword string         `json:"emphasis_keyword"`
}

type MPTemplateMsg struct {
	AppID       string                    `json:"appid"`
	TemplateID  string                    `json:"template_id"`
	Url         string                    `json:"url"`
	MiniProgram *MPTemplateMsgMiniProgram `json:"miniprogram"`
	Data        *power.HashMap            `json:"data"`
}

type MPTemplateMsgMiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}
