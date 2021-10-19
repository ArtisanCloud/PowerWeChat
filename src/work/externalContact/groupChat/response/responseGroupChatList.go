package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGroupChatList struct {
	*response.ResponseWork
	GroupChatList []*ResponseFailedChat `json:"group_chat_list"`
	NextCursor    string                `json:"next_cursor"`
}
