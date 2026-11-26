package request

type RequestWeDriveFileCreate struct {
	UserID   string `json:"userid,omitempty"`
	SpaceID  string `json:"spaceid"`
	FatherID string `json:"fatherid"`
	FileType string `json:"file_type"`
	FileName string `json:"file_name"`
}
