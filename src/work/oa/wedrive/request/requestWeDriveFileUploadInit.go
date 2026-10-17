package request

// RequestWeDriveFileUploadInit 分块上传初始化请求
type RequestWeDriveFileUploadInit struct {
	SpaceID        string   `json:"spaceid,omitempty"`         // 空间spaceid
	FatherID       string   `json:"fatherid,omitempty"`        // 当前目录的fileid，根目录时为空间spaceid
	SelectedTicket string   `json:"selected_ticket,omitempty"` // 微盘和文件选择器jsapi返回的selectedTicket。若填此参数，则不需要填spaceid/fatherid。
	FileName       string   `json:"file_name"`                 // 必填：文件名
	Size           uint64   `json:"size"`                      // 必填：文件大小（字节），最大 20G
	BlockSHA       []string `json:"block_sha"`                 // 必填：文件分块累积sha值，按分块顺序填入数组。
	SkipPushCard   bool     `json:"skip_push_card,omitempty"`  // 非必填：文件创建完成时是否推送企业微信卡片。默认false，即默认推送卡片
}
