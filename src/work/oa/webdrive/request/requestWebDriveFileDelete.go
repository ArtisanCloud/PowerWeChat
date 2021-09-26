package request

type RequestWebDriveFileDelete struct {
	UserID string `json:"userid"`
	FileID string `json:"fileid"`
}
