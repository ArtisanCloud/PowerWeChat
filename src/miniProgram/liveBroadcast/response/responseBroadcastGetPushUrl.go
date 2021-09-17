package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	*response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
