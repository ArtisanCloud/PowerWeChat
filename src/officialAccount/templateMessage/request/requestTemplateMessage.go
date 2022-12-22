package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type MiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

type RequestTemlateMessage struct {
	ToUser      string         `json:"touser"`
	TemplateID  string         `json:"template_id"`
	URL         string         `json:"url"`
	MiniProgram *MiniProgram   `json:"miniprogram"`
	Data        *power.HashMap `json:"data"`
}

type RequestTemlateMessageSubscribe struct {
	ToUser      string         `json:"touser"`
	TemplateID  string         `json:"template_id"`
	URL         string         `json:"url"`
	MiniProgram *MiniProgram   `json:"miniprogram"`
	Scene       string         `json:"scene"`
	Title       string         `json:"title"`
	Data        *power.HashMap `json:"data"`
}
