package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseAppChatGet struct {
	*response.ResponseWork

	ChatInfo *power.HashMap `json:"chat_info"`
}
