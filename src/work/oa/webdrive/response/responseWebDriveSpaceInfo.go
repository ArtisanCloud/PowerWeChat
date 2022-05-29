package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseWebDriveSpaceInfo struct {
	*response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}
