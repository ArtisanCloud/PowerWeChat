package request

type RequestWebDriveFileUpload struct {
	UserID            string `json:"userid"`
	SpaceID           string `json:"spaceid"`
	FatherID          string `json:"fatherid"`
	FileName          string `json:"file_name"`
	FileBase64Content string `json:"file_base64_content"` // 文件内容base64（注意：只需要填入文件内容的Base64，不需要添加任何如：”data:application/x-javascript;base64” 的数据类型描述信息）
}
