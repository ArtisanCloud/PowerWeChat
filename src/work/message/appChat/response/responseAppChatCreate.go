package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseAppChatCreate struct {
	*response.ResponseWork

	ChatID string `json:"chatid"`
}
