package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveSpaceShare struct {
	*response.ResponseWork

	SpaceShareURL string `json:"space_share_url"`
}
