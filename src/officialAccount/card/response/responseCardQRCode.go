package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseCreateQrCode struct {
	response.ResponseOfficialAccount

	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	URL           string `json:"url"`
	ShowQrcodeURL string `json:"show_qrcode_url"`
}
