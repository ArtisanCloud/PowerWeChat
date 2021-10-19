package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseWebDriveSpaceInfo struct {
	*response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}
