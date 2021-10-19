package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseShortLinkGenerate struct {
	*response.ResponseMiniProgram

	Link string `json:"link"`
}
