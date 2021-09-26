package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseWebDriveSpaceInfo struct {
	*response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}
