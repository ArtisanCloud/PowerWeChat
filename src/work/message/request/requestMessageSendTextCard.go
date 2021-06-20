package request

type RequestMessageSendTextCard struct {
	RequestMessageSend
	TextCard *RequestTextCard `json:"textcard"`
}

type RequestTextCard struct {
	Title       string `json:"title"`       // "领奖通知",
	Description string `json:"description"` // "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
	Url         string `json:"url"`         // "URL",
	BtnTXT      string `json:"btntxt"`      // 多"
}
