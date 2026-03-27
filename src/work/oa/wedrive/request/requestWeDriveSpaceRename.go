package request

type RequestWeDriveSpaceRename struct {
	SpaceID   string `json:"spaceid"`
	SpaceName string `json:"space_name"`
}
