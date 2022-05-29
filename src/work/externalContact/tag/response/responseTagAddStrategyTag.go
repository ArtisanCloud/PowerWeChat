package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseTagAddStrategyTag struct {
	*response.ResponseWork

	ExternalContactList *power.HashMap `json:"tag_group"`
}
