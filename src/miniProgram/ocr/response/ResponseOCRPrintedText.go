package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseOCRPrintedText struct {
	*response.ResponseMiniProgram
	Items   []*power.HashMap `json:"items"`
	ImgSize []*power.HashMap `json:"img_size"`
}
