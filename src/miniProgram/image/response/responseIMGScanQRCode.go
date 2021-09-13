package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseIMGScanQRCode struct {
	*response.ResponseMiniProgram
	CodeResults []*power.HashMap `json:"code_results"`
	ImgSize     power.HashMap    `json:"img_size"`
}
