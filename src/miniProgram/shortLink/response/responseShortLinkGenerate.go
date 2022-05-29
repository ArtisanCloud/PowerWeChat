package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseShortLinkGenerate struct {
	*response.ResponseMiniProgram

	Link string `json:"link"`
}
