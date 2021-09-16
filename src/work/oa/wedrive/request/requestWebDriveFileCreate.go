package request

type RequestWebDriveFileCreate struct {
	UserID   string `json:"userid"`
	SpaceID  string `json:"spaceid"`
	FatherID string `json:"fatherid"`
	FileType string `json:"file_type"`
	FileName string `json:"file_name"`
}
