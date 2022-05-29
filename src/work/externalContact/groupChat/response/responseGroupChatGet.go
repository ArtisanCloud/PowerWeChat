package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGroupChatGet struct {
	*response.ResponseWork
	GroupChat *power.HashMap `json:"group_chat"`
}
