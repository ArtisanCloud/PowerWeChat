package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	*response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`

}
