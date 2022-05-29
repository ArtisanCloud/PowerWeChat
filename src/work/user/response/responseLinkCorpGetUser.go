package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	*response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`
}
