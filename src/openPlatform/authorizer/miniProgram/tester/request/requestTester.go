package request

type RequestUnbind struct {
	UserStr  string `json:"userstr"`
	wechatID string `json:"wechatid"`
}
