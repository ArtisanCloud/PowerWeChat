package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseShortLinkGenerate struct {
	*response.ResponseMiniProgram

	Link string `json:"link"`
}
