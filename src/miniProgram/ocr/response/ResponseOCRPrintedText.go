package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseOCRPrintedText struct {
	*response.ResponseMiniProgram
	Items   []*power.HashMap `json:"items"`
	ImgSize []*power.HashMap `json:"img_size"`
}
