package request

type RequestWeDriveFileUpload struct {
	SpaceID           string `json:"spaceid"`             // 空间spaceid
	FatherID          string `json:"fatherid"`            // 父目录fileid, 在根目录时为空间spaceid
	SelectedTicket    string `json:"selected_ticket"`     // 微盘和文件选择器jsapi返回的selectedTicket。若填此参数，则不需要填spaceid/fatherid。
	FileName          string `json:"file_name"`           // 文件名字（注意：文件名最多填255个字符, 英文算1个, 汉字算2个）
	FileBase64Content string `json:"file_base64_content"` // 文件内容base64（注意：只需要填入文件内容的Base64，不需要添加任何如："data:application/x-javascript;base64" 的数据类型描述信息），文件大小上限为10M。大于10M文件，可使用文件分块上传接口
}
