package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCustomerGetUpgradeServiceConfig struct {
	*response.ResponseWork

	MemberRange    *power.HashMap `json:"member_range"`
	GroupChatRange *power.HashMap `json:"groupchat_range"`
}
