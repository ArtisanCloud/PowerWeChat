package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGroupChatGet struct {
	*response.ResponseWork
	GroupChat *power.HashMap `json:"group_chat"`
}
