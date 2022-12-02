package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseGettingRulesForAutoReplies struct {
	*response.ResponseOfficialAccount

	IsAddFriendReplyOpen        int                          `json:"is_add_friend_reply_open"`
	IsAutoreplyOpen             int                          `json:"is_autoreply_open"`
	AddFriendAutoReplyInfo      *AddFriendAutoReplyInfo      `json:"add_friend_autoreply_info"`
	MessageDefaultAutoReplyInfo *MessageDefaultAutoReplyInfo `json:"message_default_autoreply_info"`
	KeywordAutoReplyInfo        *KeywordAutoReplyInfo        `json:"keyword_autoreply_info"`
}

type AddFriendAutoReplyInfo struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type MessageDefaultAutoReplyInfo struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type KeywordInfo struct {
	Type      string `json:"type"`
	MatchMode string `json:"match_mode"`
	Content   string `json:"content"`
}

type NewsInfo struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Digest     string `json:"digest"`
	ShowCover  int    `json:"show_cover"`
	CoverUrl   string `json:"cover_url"`
	ContentUrl string `json:"content_url"`
	SourceUrl  string `json:"source_url"`
}

type NewsInfos struct {
	List []NewsInfo `json:"list"`
}

type ReplyInfo struct {
	Type      string     `json:"type"`
	NewsInfos *NewsInfos `json:"news_info,omitempty"`
	Content   string     `json:"content,omitempty"`
}

type AutoReplyInfo struct {
	RuleName        string         `json:"rule_name"`
	CreateTime      int            `json:"create_time"`
	ReplyMode       string         `json:"reply_mode"`
	KeywordListInfo []*KeywordInfo `json:"keyword_list_info"`
	ReplyListInfo   []*ReplyInfo   `json:"reply_list_info"`
}

type KeywordAutoReplyInfo struct {
	List []*AutoReplyInfo `json:"list"`
}
