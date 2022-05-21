package request

type RequestMessageSendNews struct {
	RequestMessageSend
	News *RequestNews `json:"news"`
}

type RequestNews struct {
	Article []*RequestNewsArticle `json:"articles"`
}

type RequestNewsArticle struct {
	Title       string `json:"title"`       // "领奖通知",
	Description string `json:"description"` // "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
	URL         string `json:"url"`         // "URL",
	PicURL      string `json:"picurl"`      // 多"
	AppID       string `json:"appid"`
	PagePath    string `json:"pagepath"`
}
