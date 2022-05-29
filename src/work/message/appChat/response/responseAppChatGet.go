package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseAppChatGet struct {
	*response.ResponseWork

	ChatInfo *power.HashMap `json:"chat_info"`
}
