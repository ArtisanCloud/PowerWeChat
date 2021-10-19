package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseTagAddStrategyTag struct {
	*response.ResponseWork

	ExternalContactList *power.HashMap `json:"tag_group"`
}
