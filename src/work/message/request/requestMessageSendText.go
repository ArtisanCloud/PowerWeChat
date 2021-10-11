package request

type RequestMessageSendText struct {
	RequestMessageSend
	Text *RequestText `json:"text"`
}

type RequestText struct {
	Content string `json:"content"` // {"你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。"},
}
