package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseURLLinkGenerate struct {
	*response.ResponseMiniProgram
	URLLink string `json:"url_link"`
}
