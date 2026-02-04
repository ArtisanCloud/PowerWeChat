package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseWeDriveFileUploadFinish 分块上传完成响应
type ResponseWeDriveFileUploadFinish struct {
	response.ResponseWork
	UploadKey string `json:"upload_key"` // 文件上传凭证。file_upload_init返回的upload_key
}
