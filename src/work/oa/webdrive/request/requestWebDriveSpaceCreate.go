package request

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"

type RequestWebDriveSpaceCreate struct {
	UserID    string           `json:"userid"`
	SpaceName string           `json:"space_name"`
	AuthInfo  []*power.HashMap `json:"auth_info"`
}
