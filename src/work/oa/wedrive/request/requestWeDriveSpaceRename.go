package request

type RequestWeDriveSpaceRename struct {
	UserID    string `json:"userid,omitempty"`
	SpaceID   string `json:"spaceid"`
	SpaceName string `json:"space_name"`
}
