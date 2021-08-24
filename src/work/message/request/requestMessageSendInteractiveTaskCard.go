package request

type RequestMessageSendInteractiveTaskCard struct {
	RequestMessageSend
	InteractiveTaskCard *InteractiveTaskCard `json:"interactive_taskcard"`
}

type InteractiveTaskCard struct {
	Title       string `json:"title"`       //  "会议室预订成功通知",
	Description string `json:"description"` //  "4月27日 16:16",
	URl         string `json:"url"`         //  true,
	BTN         []*BTN `json:"content_item"`
}

type BTN struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	IsBold bool   `json:"is_bold"`
}
