package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseIMGSuperResolution struct {
	*response.ResponseMiniProgram
	MediaID string `json:"media_id"`
}
