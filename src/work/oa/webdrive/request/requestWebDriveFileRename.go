package request

type RequestWebDriveFileRename struct {
	UserID  string `json:"userid"`
	FileID  string `json:"fileid"`
	NewName string `json:"new_name"`
}
