package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseExpressAccountGetAll struct {
	*response2.ResponseMiniProgram

	Count int               `json:"count"`
	List  []*power.HashMap `json:"list"`
}
