package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseIMGScanQRCode struct {
	*response.ResponseMiniProgram
	CodeResults []*power.HashMap `json:"code_results"`
	ImgSize     power.HashMap    `json:"img_size"`
}
