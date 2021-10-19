package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	*response.ResponseWork
	ChatID *power.HashMap `json:"chat_id"`
}
