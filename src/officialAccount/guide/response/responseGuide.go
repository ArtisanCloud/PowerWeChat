package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type GuideAdviser struct {
	GuideAccount    string `json:"guide_account"`
	GuideHeadImgURL string `json:"guide_headimgurl"`
	GuideNickname   string `json:"guide_nickname"`
	GuideOpenid     int    `json:"guide_openid"`
}

type ResponseGuideGetAdviser struct {
	*response.ResponseOfficialAccount

	*GuideAdviser
}

// ------------------------------------------------------------

type ResponseGuideGetAdvisers struct {
	*response.ResponseOfficialAccount

	TotalNum int             `json:"total_num"`
	List     []*GuideAdviser `json:"list"`
}

// ------------------------------------------------------------

type Message struct {
	GuideAccount string `json:"guide_account"`
	Openid       string `json:"openid"`
	CreateTime   int    `json:"create_time"`
	Content      string `json:"content"`
	ContentType  int    `json:"content_type"`
	Direction    int    `json:"direction"`
}

type ResponseGuideGetChatRecords struct {
	*response.ResponseOfficialAccount

	TotalNum int        `json:"total_num"`
	MsgList  []*Message `json:"msg_list"`
}

// ------------------------------------------------------------

type AutoReply struct {
	MsgType    int    `json:"msgtype"`
	Content    string `json:"content"`
	UpdateTime int    `json:"updatetime"`
}

type ResponseGuideGetConfig struct {
	*response.ResponseOfficialAccount

	GuideFastReplyList []interface{} `json:"guide_fast_reply_list"`
	GuideAutoReply     *AutoReply    `json:"guide_auto_reply"`
	UpdateTime         int           `json:"updatetime"`
	GuideAutoReplyPlus *AutoReply    `json:"guide_auto_reply_plus"`
}

// ------------------------------------------------------------

type BlackKeyword struct {
	Values     []string `json:"values"`
	Updatetime int      `json:"updatetime"`
}

type ResponseGuideGetAdviserConfig struct {
	*response.ResponseOfficialAccount

	BlackKeyword struct {
		Values         []*BlackKeyword `json:"black_keyword"`
		GuideAutoReply AutoReply       `json:"guide_auto_reply"`
	}
}

// ------------------------------------------------------------

type ResponseGuideCreateGroup struct {
	*response.ResponseOfficialAccount

	GroupID interface{} `json:"group_id"`
}

// ------------------------------------------------------------

type Group struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}

type ResponseGuideGetGroupList struct {
	*response.ResponseOfficialAccount

	GroupList []*Group `json:"group_list"`
}

// ------------------------------------------------------------

type Guide struct {
	GuideAccount    string `json:"guide_account"`
	GuideHeadImgURL string `json:"guide_headimgurl"`
	GuideNickname   string `json:"guide_nickname"`
	CreateTime      int    `json:"create_time"`
	GuideOpenid     string `json:"guide_openid"`
}
type ResponseGuideGetGroups struct {
	*response.ResponseOfficialAccount

	GuideList []*Guide `json:"guide_list"`
	TotalNum  int      `json:"total_num"`
}

// ------------------------------------------------------------

type ResponseGuideGetGuideGroup struct {
	*response.ResponseOfficialAccount

	GroupIDList []int `json:"group_id_list"`
}

// ------------------------------------------------------------

type BuyerResp struct {
	*response.ResponseOfficialAccount

	OpenID string `json:"openid"`
}

type ResponseGuideBuyerRelation struct {
	*response.ResponseOfficialAccount

	BuyerResp []*BuyerResp `json:"buyer_resp"`
}

type BuyerRelation struct {
	OpenID        string `json:"openid"`
	GuideAccount  string `json:"guide_account"`
	GuideOpenID   string `json:"guide_openid"`
	BuyerNickName string `json:"buyer_nickname"`
	CreateTime    int    `json:"create_time"`
}

type ResponseGuideBuyerRelationList struct {
	*response.ResponseOfficialAccount

	TotalNum int             `json:"total_num"`
	List     []BuyerRelation `json:"list"`
}

type ResponseGuideGetBuyerRelation struct {
	*response.ResponseOfficialAccount

	*BuyerRelation
}

// ------------------------------------------------------------

type Option struct {
	TagName   string   `json:"tag_name"`
	TagValues []string `json:"tag_values"`
}

type ResponseGuideTagOption struct {
	*response.ResponseOfficialAccount

	Options []*Option `json:"options"`
}

// ------------------------------------------------------------

type ResponseGuideGetBuyerTags struct {
	*response.ResponseOfficialAccount

	TagValues []string `json:"tag_values"`
}

// ------------------------------------------------------------

type ResponseGuideGetBuyerByTag struct {
	*response.ResponseOfficialAccount

	OpenidList []string `json:"openid_list"`
}

// ------------------------------------------------------------

type ResponseGuideGetBuyerDisplayTags struct {
	*response.ResponseOfficialAccount

	DisplayTagList []string `json:"display_tag_list"`
}

// ------------------------------------------------------------

type Card struct {
	Title    string      `json:"title"`
	AppID    string      `json:"appid"`
	Path     string      `json:"path"`
	PicURL   string      `json:"picurl"`
	MasterID interface{} `json:"master_id"`
	SlaveID  interface{} `json:"slave_id"`
}

type ResponseGuideGetCardMaterial struct {
	*response.ResponseOfficialAccount

	CardList []*Card `json:"card_list"`
}

// ------------------------------------------------------------

type Model struct {
	PicURL string `json:"picurl"`
}

type ResponseGuideGetImageMaterial struct {
	*response.ResponseOfficialAccount

	ModelList []*Model `json:"model_list"`
	TotalNum  int      `json:"total_num"`
}

// ------------------------------------------------------------

type Word struct {
	Word       string `json:"word"`
	CreateTime int    `json:"create_time"`
}

type ResponseGuideGetWordMaterial struct {
	*response.ResponseOfficialAccount

	WordList []*Word `json:"word_list"`
	TotalNum int     `json:"total_num"`
}
