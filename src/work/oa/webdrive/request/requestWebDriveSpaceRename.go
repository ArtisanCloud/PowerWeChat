package request

type RequestWebDriveSpaceRename struct {
	UserID    string `json:"userid"`
	SpaceID   string `json:"spaceid"`
	SpaceName string `json:"space_name"`
}
