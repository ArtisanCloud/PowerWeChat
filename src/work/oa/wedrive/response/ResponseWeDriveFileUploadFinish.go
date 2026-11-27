package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseWeDriveFileUploadFinish 分块上传完成响应
type ResponseWeDriveFileUploadFinish struct {
	response.ResponseWork
	Fileid string `json:"fileid"` // 文件fileid
}
