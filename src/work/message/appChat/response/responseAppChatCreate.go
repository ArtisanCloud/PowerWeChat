package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseAppChatCreate struct {
	*response.ResponseWork

	ChatID string `json:"chatid"`
}
