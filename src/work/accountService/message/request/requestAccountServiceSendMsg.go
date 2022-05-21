package request

type RequestAccountServiceSendMsg struct {
	ToUser      string                              `json:"touser"`
	OpenKfid    string                              `json:"open_kfid"`
	MsgID       string                              `json:"msgid"`
	MsgType     string                              `json:"msgtype"`
	Text        RequestAccountServiceMsgText        `json:"text,omitempty"`
	Image       RequestAccountServiceMsgImage       `json:"image,omitempty"`
	Voice       RequestAccountServiceMsgVoice       `json:"voice,omitempty"`
	File        RequestAccountServiceMsgFile        `json:"file,omitempty"`
	Link        RequestAccountServiceMsgLink        `json:"link,omitempty"`
	MiniProgram RequestAccountServiceMsgMiniProgram `json:"miniprogram,omitempty"`
	Menu        RequestAccountServiceMsgMenu        `json:"msgmenu,omitempty"`
	Location    RequestAccountServiceMsgLocation    `json:"location,omitempty"`
}

type RequestAccountServiceMsgText struct {
	Content string `json:"content"`
}

type RequestAccountServiceMsgImage struct {
	MediaID string `json:"media_id"`
}

type RequestAccountServiceMsgVoice struct {
	MediaID string `json:"media_id"`
}

type RequestAccountServiceMsgFile struct {
	MediaID string `json:"media_id"`
}

type RequestAccountServiceMsgLink struct {
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Url          string `json:"url"`
	ThumbMediaID string `json:"thumb_media_id"`
}

type RequestAccountServiceMsgMiniProgram struct {
	AppID        string `json:"appid"`
	Title        string `json:"title"`
	ThumbMediaID string `json:"thumb_media_id"`
	PagePath     string `json:"pagepath"`
}

type RequestAccountServiceMsgMenu struct {
	HeadContent string                             `json:"head_content"`
	TailContent string                             `json:"tail_content,omitempty"`
	List        []RequestAccountServiceMsgMenuList `json:"list,omitempty"`
}
type RequestAccountServiceMsgMenuList struct {
	Click       RequestAccountServiceMsgMenuListClick       `json:"click,omitempty"`
	View        RequestAccountServiceMsgMenuListView        `json:"view,omitempty"`
	MiniProgram RequestAccountServiceMsgMenuListMiniProgram `json:"miniprogram,omitempty"`
}
type RequestAccountServiceMsgMenuListClick struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
type RequestAccountServiceMsgMenuListView struct {
	Url     string `json:"url"`
	Content string `json:"content"`
}
type RequestAccountServiceMsgMenuListMiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Content  string `json:"content"`
}

type RequestAccountServiceMsgLocation struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}
