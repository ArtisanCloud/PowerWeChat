package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	response.ResponseWork
	UserInfo *power.HashMap `json:"user_info"`

}
