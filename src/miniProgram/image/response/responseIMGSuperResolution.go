package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseIMGSuperResolution struct {
	*response.ResponseMiniProgram
	MediaID string `json:"media_id"`
}
