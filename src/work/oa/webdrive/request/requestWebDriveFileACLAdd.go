package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestWebDriveFileACLAdd struct {
	UserID   string           `json:"userid"`
	FileID   string           `json:"fileid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}
