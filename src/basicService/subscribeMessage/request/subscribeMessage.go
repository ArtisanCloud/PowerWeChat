package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestSubscribeMessageSend struct {
	ToUser           string         `json:"touser"`
	TemplateID       string         `json:"template_id"`
	Page             string         `json:"page,omitempty"`
	MiniProgramState string         `json:"miniprogram_state,omitempty"`
	Lang             string         `json:"lang,omitempty"`
	Data             *power.HashMap `json:"data"`
}

type MiniProgram struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type RequestSubscribeMessageBizSend struct {
	ToUser      string         `json:"touser"`
	TemplateId  string         `json:"template_id"`
	Page        string         `json:"page"`
	MiniProgram *MiniProgram   `json:"miniprogram"`
	Data        *power.HashMap `json:"data"`
}
