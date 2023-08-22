package response

type UploadItem struct {
	Id           string `json:"id"`
	ItemUrl      string `json:"item_url"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	UploadStatus int    `json:"upload_status"`
	ErrMsg       string `json:"errmsg"` // 这个upload status与官方文档返回不一致 可以直接判断errmsg是否为空 代表是否有错误
}

type UploadProductSearchResponse struct {
	Cost       int           `json:"cost"`
	End        int           `json:"end"`
	Errcode    int           `json:"errcode"`
	Errmsg     string        `json:"errmsg"`
	Progress   int           `json:"progress"`
	Status     int           `json:"status"`
	Total      int           `json:"total"`
	UploadItem []*UploadItem `json:"upload_item"`
}
