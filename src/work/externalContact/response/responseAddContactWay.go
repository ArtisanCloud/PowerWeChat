package response

import "github.com/ArtisanCloud/power-wechat/src/kernel/response"

type ResponseAddContactWay struct {
	*response.ResponseWork
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}
