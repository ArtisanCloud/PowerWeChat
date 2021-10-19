package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCustomerGetUpgradeServiceConfig struct {
	*response.ResponseWork

	MemberRange    *power.HashMap `json:"member_range"`
	GroupChatRange *power.HashMap `json:"groupchat_range"`
}
