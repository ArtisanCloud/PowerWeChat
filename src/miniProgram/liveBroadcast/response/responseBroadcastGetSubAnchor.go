package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	*response.ResponseMiniProgram

	UserName int `json:"username"`
}
