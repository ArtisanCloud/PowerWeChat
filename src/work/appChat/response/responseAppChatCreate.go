package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseAppChatCreate struct {
	*response.ResponseWork

	ChatID string `json:"chatid"`
}
