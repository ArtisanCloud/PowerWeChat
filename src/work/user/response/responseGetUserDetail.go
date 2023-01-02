package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetUserDetail struct {
	response.ResponseWork

	*UserDetail
}

type ExtAttr struct {
	Attrs []struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Text struct {
			Value string `json:"value"`
		} `json:"text,omitempty"`
		Web struct {
			Url   string `json:"url"`
			Title string `json:"title"`
		} `json:"web,omitempty"`
	} `json:"attrs"`
}

type WechatChannels struct {
	NickName string `json:"nickname"`
	Status   int    `json:"status"`
}

type Text struct {
	Value string `json:"value"`
}

type Web struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type MiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Title    string `json:"title"`
}

type ExternalAttr struct {
	Type        int          `json:"type"`
	Name        string       `json:"name"`
	Text        *Text        `json:"text,omitempty"`
	Web         *Web         `json:"web,omitempty"`
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"`
}

type ExternalProfile struct {
	ExternalCorpName string          `json:"external_corp_name"`
	WechatChannels   *WechatChannels `json:"wechat_channels"`
	ExternalAttr     []*ExternalAttr `json:"external_attr"`
}

type UserDetail struct {
	UserID           string           `json:"userid"`
	Name             string           `json:"name"`
	Department       []int            `json:"department"`
	Order            []int            `json:"order"`
	Position         string           `json:"position"`
	Mobile           string           `json:"mobile"`
	Gender           string           `json:"gender"`
	Email            string           `json:"email"`
	BizMail          string           `json:"biz_mail"`
	IsLeaderInDept   []int            `json:"is_leader_in_dept"`
	DirectLeader     []string         `json:"direct_leader"`
	Avatar           string           `json:"avatar"`
	ThumbAvatar      string           `json:"thumb_avatar"`
	Telephone        string           `json:"telephone"`
	Alias            string           `json:"alias"`
	Address          string           `json:"address"`
	OpenUserid       string           `json:"open_userid"`
	MainDepartment   int              `json:"main_department"`
	ExtAttr          *ExtAttr         `json:"extattr"`
	Status           int              `json:"status"`
	QrCode           string           `json:"qr_code"`
	ExternalPosition string           `json:"external_position"`
	ExternalProfile  *ExternalProfile `json:"external_profile"`
}
