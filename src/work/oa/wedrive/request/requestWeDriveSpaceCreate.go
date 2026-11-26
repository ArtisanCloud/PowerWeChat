package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestWeDriveSpaceCreate struct {
	UserID       string           `json:"userid,omitempty"`
	SpaceSubType uint32           `json:"space_sub_type,omitempty"`
	SpaceName    string           `json:"space_name"`
	AuthInfo     []*power.HashMap `json:"auth_info"`
}
