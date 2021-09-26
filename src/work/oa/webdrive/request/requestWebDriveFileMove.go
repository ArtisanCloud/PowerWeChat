package request

type RequestWebDriveFileMove struct {
	UserID   string   `json:"userid"`
	FatherID string   `json:"fatherid"`
	Replace  bool     `json:"replace"`
	FileID   []string `json:"fileid"`
}
