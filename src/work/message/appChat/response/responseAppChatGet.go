package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAppChatGet struct {
	*response.ResponseWork

	ChatInfo *power.HashMap `json:"chat_info"`
}
