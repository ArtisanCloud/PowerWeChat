package request

type RequestWebDriveFileSetting struct {
	UserID    string `json:"userid"`
	FileID    string `json:"fileid"`
	AuthScope int    `json:"auth_scope"`
	Auth      int    `json:"auth"`
}
