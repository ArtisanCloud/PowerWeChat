package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGroupChatTransfer struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"failed_chat_list"`
}
