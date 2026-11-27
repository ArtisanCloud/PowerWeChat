package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseWeDriveFileUploadInit 分块上传初始化响应
type ResponseWeDriveFileUploadInit struct {
	response.ResponseWork
	UploadKey string `json:"upload_key"` // 文件上传凭证。不命中秒传时返回，作为file_upload_part参数
	FileID    string `json:"fileid"`     // 文件fileid。命中秒传时返回，此时上传流程完成
	HitExist  bool   `json:"hit_exist"`  // 是否命中秒传
}
