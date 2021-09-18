package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGroupChatTransfer struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"failed_chat_list"`
}
