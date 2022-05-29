package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	*response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
