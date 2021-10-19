package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseWebDriveSpaceInfo struct {
	*response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}
