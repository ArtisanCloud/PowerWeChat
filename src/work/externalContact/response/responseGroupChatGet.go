package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGroupChatGet struct {
	response.ResponseWork
	GroupChat *power.HashMap `json:"group_chat"`
}
