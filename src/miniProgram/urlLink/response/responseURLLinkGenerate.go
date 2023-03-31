package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseURLLinkGenerate struct {
	response.ResponseMiniProgram
	URLLink string `json:"url_link"`
}
