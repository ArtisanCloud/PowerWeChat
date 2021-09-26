package request

type RequestWebDriveFileList struct {
	UserID   string `json:"userid"`
	SpaceID  string `json:"spaceid"`
	FatherID string `json:"fatherid"`
	SortType int    `json:"sort_type"`
	Start    int    `json:"start"`
	Limit    int    `json:"limit"`
}
