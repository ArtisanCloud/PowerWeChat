package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAppChatGet struct {
	*response.ResponseWork

	ChatInfo *power.HashMap `json:"chat_info"`
}
