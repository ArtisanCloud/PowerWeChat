package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type CompactGroupChat struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

type ResponseGroupChatList struct {
	response.ResponseWork
	GroupChatList []*CompactGroupChat `json:"group_chat_list"`
	NextCursor    string              `json:"next_cursor"`
}
