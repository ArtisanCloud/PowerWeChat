package response

import "github.com/ArtisanCloud/go-wechat/src/kernel/response"

type ResponseAddContactWay struct {
	*response.ResponseWX
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}
