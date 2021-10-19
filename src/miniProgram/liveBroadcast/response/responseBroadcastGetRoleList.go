package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseBroadcastGetRoleList struct {
	*response.ResponseMiniProgram

	Total int              `json:"total,omitempty"`
	List  []*power.HashMap `json:"list,omitempty"`
}
