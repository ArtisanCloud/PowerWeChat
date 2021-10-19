package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGroupChatTransfer struct {
	*response.ResponseWork

	Customer []*power.HashMap `json:"failed_chat_list"`
}
