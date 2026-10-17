package request

// RequestWeDriveFileUploadFinish 分块上传完成请求
type RequestWeDriveFileUploadFinish struct {
	UploadKey string `json:"upload_key"` // 必填：init 接口返回的 upload_key
}
