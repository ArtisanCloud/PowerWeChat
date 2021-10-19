package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseBroadcastGetPushUrl struct {
	*response.ResponseMiniProgram

	PushAddr string `json:"pushAddr"`
}
