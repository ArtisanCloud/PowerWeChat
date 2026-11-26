package request

// RequestWeDriveFileUploadPart 分块上传单个文件块
type RequestWeDriveFileUploadPart struct {
	UploadKey         string `json:"upload_key"`          // 必填：文件上传凭证。file_upload_init返回的upload_key
	Index             int32  `json:"index"`               // 必填：文件分块号。文件内容按2M分块，从1开始
	FileBase64Content string `json:"file_base64_content"` // 必填：分块的文件内容base64。（注意：只需要填入文件内容的Base64，不需要添加任何如："data:application/x-javascript;base64" 的数据类型描述信息）
}
