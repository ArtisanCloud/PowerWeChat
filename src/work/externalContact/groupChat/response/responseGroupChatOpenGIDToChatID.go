package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	*response.ResponseWork
	ChatID *power.HashMap `json:"chat_id"`
}
