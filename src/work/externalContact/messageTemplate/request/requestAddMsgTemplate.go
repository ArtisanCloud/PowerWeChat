package request

type GroupList struct {
	TagList []string `json:"tag_list"`
}

type TagFilter struct {
	GroupList []GroupList `json:"group_list"`
}

type RequestAddMsgTemplate struct {
	ChatType       string                     `json:"chat_type"`
	ExternalUserID []string                   `json:"external_userid"`
	ChatIdList     []string                   `json:"chat_id_list"`
	TagFilter      TagFilter                  `json:"tag_filter,omitempty"`
	Sender         string                     `json:"sender,omitempty"`
	AllowSelect    bool                       `json:"allow_select"`
	Text           *TextOfMessage             `json:"text,omitempty"`
	Attachments    []AttachmentOfMessage      `json:"attachments,omitempty"`
	Image          *ImageMediaOfMessage       `json:"image,omitempty"`
	Link           *LinkMsgOfMessage          `json:"link,omitempty"`
	MiniProgram    *MiniProgramMediaOfMessage `json:"miniprogram,omitempty"`
	Video          *VideoMediaOfMessage       `json:"video,omitempty"`
	File           *FileMediaOfMessage        `json:"file,omitempty"`
}

type AttachmentOfMessage struct {
	MsgType string `json:"msgtype"`
}

type ImageMediaOfMessage struct {
	MediaId string `json:"media_id"`
	PicURL  string `json:"pic_url"`
}

type LinkMsgOfMessage struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}
type MiniProgramMediaOfMessage struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}
type VideoMediaOfMessage struct {
	MediaId string `json:"media_id"`
}
type FileMediaOfMessage struct {
	MediaId string `json:"media_id"`
}
