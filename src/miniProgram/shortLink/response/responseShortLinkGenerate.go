package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseShortLinkGenerate struct {
	*response.ResponseMiniProgram

	Link string `json:"link"`
}
