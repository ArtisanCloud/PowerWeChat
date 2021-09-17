package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGetRoleList struct {
	*response.ResponseMiniProgram

	Total int              `json:"total"`
	List  []*power.HashMap `json:"list"`
}
