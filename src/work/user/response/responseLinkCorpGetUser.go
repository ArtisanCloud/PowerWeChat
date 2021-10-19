package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	*response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`

}
