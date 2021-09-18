package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseTagAddCorpTag struct {
	*response.ResponseWork

	ExternalContactList *power.HashMap `json:"tag_group"`
}
