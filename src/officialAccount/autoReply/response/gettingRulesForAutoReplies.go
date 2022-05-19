package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ResponseGettingRulesForAutoReplies struct {
	*response.ResponseOfficialAccount

	IsAddFriendReplyOpen        int                          `json:"is_add_friend_reply_open"`
	IsAutoReplyOpen             int                          `json:"is_autoreply_open"`
	AddFriendAutoReplyInfo      *AddFriendAutoReplyInfo      `json:"add_friend_autoreply_info"`
	MessageDefaultAutoReplyInfo *MessageDefaultAutoReplyInfo `json:"message_default_autoreply_info"`
	KeywordAutoReplyInfo        []*KeywordAutoReplyInfo      `json:"keyword_autoreply_info"`
}

type AddFriendAutoReplyInfo struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type MessageDefaultAutoReplyInfo struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type KeywordAutoReplyInfo struct {
	List []interface{} `json:"list"`
}

type ReplyInfo struct {
}
