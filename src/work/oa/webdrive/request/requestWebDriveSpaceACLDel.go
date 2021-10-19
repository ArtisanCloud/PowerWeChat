package request

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"

type RequestWebDriveSpaceACLDel struct {
	UserID    string `json:"userid"`
	SpaceID   string `json:"spaceid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}
