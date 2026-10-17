package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestWeDriveSpaceCreate struct {
	SpaceSubType uint32           `json:"space_sub_type"`
	SpaceName    string           `json:"space_name"`
	AuthInfo     []*power.HashMap `json:"auth_info"`
}
