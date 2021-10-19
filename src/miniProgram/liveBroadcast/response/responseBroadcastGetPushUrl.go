package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	*response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
