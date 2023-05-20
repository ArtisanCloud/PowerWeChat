package request

type TextOfMessage struct {
	Content string `json:"content"` // "content",

}

type ImageOfMessage struct {
	MsgType string `json:"msgtype"` // "image"
	Image   *Image `json:"image"`
}

func (media *ImageOfMessage) GetMsgType() string {
	return "image"
}

type Image struct {
	MediaID string `json:"media_id"` // "MEDIA_ID",
	PicURL  string `json:"pic_url"`  // "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0"
}

type LinkOfMessage struct {
	MsgType string `json:"msgtype"` // "link"
	Link    *Link  `json:"link"`
}

func (media *LinkOfMessage) GetMsgType() string {
	return "link"
}

type Link struct {
	Title  string `json:"title"`  //  "消息标题",
	PicURL string `json:"picurl"` //  "https://example.pic.com/path",
	Desc   string `json:"desc"`   //  "消息描述",
	URL    string `json:"url"`    //  "https://example.link.com/path"
}

type MiniProgramOfMessage struct {
	MsgType     string       `json:"msgtype"` // "miniprogram"
	MiniProgram *MiniProgram `json:"miniprogram"`
}

func (media *MiniProgramOfMessage) GetMsgType() string {
	return "miniprogram"
}

type MiniProgram struct {
	Title      string `json:"title"`        //  "消息标题",
	PicMediaID string `json:"pic_media_id"` //  "MEDIA_ID",
	AppID      string `json:"appid"`        //  "wx8bd80126147dfAAA",
	Page       string `json:"page"`         //  "/path/index.html"
}

type File struct {
	MediaID string `json:"media_id"`
}
type FileOfMessage struct {
	MsgType string `json:"msgtype"` // "file"
	File    *File  `json:"file"`
}

func (media *FileOfMessage) GetMsgType() string {
	return "file"
}

type VideoOfMessage struct {
	MsgType string `json:"msgtype"` // "video"
	Video   *Video `json:"video"`
}

func (media *VideoOfMessage) GetMsgType() string {
	return "video"
}

type Video struct {
	MediaID string `json:"media_id"` // "MEDIA_ID",
}

type MessageTemplateInterface interface {
	GetMsgType() string
}

type Attachment struct {
	MsgType     string       `json:"msgtype"`
	Image       *Image       `json:"image,omitempty"`
	Link        *Link        `json:"link,omitempty"`
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"`
	Video       *Video       `json:"video,omitempty"`
	File        *File        `json:"file,omitempty"`
}

func (attachment *Attachment) GetMsgType() string {
	return attachment.MsgType
}
