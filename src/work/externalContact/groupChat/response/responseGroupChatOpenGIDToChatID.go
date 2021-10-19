package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	*response.ResponseWork
	ChatID *power.HashMap `json:"chat_id"`
}
