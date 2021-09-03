package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseURLLinkGenerate struct {
	*response.ResponseMiniProgram
	URLLink string `json:"url_link"`
}
