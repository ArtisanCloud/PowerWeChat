package request

type TextOfMessage struct {
	Content string // "content",

}

type ImageOfMessage struct {
	MsgType string // "image"
	Image   Image
}

type Image struct {
	MediaID string // "MEDIA_ID",
	PicURL  string // "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0"
}

type LinkOfMessage struct {
	MsgType string // "link"
	Link    Link
}

type Link struct {
	Title  string //  "消息标题",
	PicURL string //  "https://example.pic.com/path",
	Desc   string //  "消息描述",
	URL    string //  "https://example.link.com/path"
}

type MiniProgramOfMessage struct {
	MsgType     string // "miniprogram"
	MiniProgram MiniProgram
}

type MiniProgram struct {
	Title      string //  "消息标题",
	PicMediaID string //  "MEDIA_ID",
	AppID      string //  "wx8bd80126147dfAAA",
	Page       string //  "/path/index.html"
}

type VideoOfMessage struct {
	MsgType string // "video"
	Video   Video
}

type Video struct {
	MediaID string // "MEDIA_ID",
}
