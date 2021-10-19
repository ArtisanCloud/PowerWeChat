package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseGroupChatGet struct {
	*response.ResponseWork
	GroupChat *power.HashMap `json:"group_chat"`
}
