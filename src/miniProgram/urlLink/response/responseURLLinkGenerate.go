package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseURLLinkGenerate struct {
	*response.ResponseMiniProgram
	URLLink string `json:"url_link"`
}
