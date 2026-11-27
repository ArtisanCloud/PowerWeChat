package request

type RequestWeDriveFileDelete struct {
	UserID string `json:"userid"`
	FileID string `json:"fileid"`
}

// RequestWeDriveFileDeleteBatch deletes multiple files in one request.
type RequestWeDriveFileDeleteBatch struct {
	UserID  string   `json:"userid,omitempty"`
	FileIDs []string `json:"fileid"`
}
