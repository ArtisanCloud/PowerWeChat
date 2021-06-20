package request

type RequestMessageSendMiniProgramNotice struct {
	RequestMessageSend
	MiniProgramNotice *MiniProgramNotice `json:"miniprogram_notice"`
}

type MiniProgramNotice struct {
	Appid             string         `json:"appid"`               //  "wx123123123123123",
	Page              string         `json:"page"`                //  "pages/index?userid=zhangsan&orderid=123123123",
	Title             string         `json:"title"`               //  "会议室预订成功通知",
	Description       string         `json:"description"`         //  "4月27日 16:16",
	EmphasisFirstItem bool           `json:"emphasis_first_item"` //  true,
	ContentItem       []*ContentItem `json:"content_item"`
}

type ContentItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
