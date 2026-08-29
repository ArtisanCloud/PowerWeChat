package request

type RequestWeDriveFileDownload struct {
	FileID         string `json:"fileid,omitempty"`          // 文件fileid（只支持下载普通文件，不支持下载文件夹或微文档）
	SelectedTicket string `json:"selected_ticket,omitempty"` // 微盘和文件选择器jsapi返回的selectedTicket。若填此参数，则不需要填fileid。
}
