package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Text struct {
	Content string `json:"content"`
}

type Image struct {
	PicURL string `json:"pic_url"`
}

type Link struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}

type MiniProgram struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}

type Conclusions struct {
	Text        *Text        `json:"text"`
	Image       *Image       `json:"image"`
	Link        *Link        `json:"link"`
	MiniProgram *MiniProgram `json:"miniprogram"`
}

type ContactWay struct {
	ConfigID string `json:"config_id"`

	Type          int          `json:"type"`
	Scene         int          `json:"scene"`
	Style         int          `json:"style"`
	Remark        string       `json:"remark"`
	SkipVerify    bool         `json:"skip_verify"`
	State         string       `json:"state"`
	QrCode        string       `json:"qr_code"`
	User          []string     `json:"user"`
	Party         []int        `json:"party"`
	IsTemp        bool         `json:"is_temp"`
	ExpiresIn     int          `json:"expires_in"`
	ChatExpiresIn int          `json:"chat_expires_in"`
	UnionID       string       `json:"unionid"`
	Conclusions   *Conclusions `json:"conclusions"`
}

type ResponseGetContactWay struct {
	response.ResponseWork

	ContactWay *ContactWay `json:"contact_way"`
}
