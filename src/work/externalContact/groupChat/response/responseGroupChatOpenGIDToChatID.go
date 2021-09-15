package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	*response.ResponseWork
	ChatID *power.HashMap `json:"chat_id"`
}
