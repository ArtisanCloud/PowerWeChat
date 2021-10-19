package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	*response.ResponseMiniProgram

	UserName string `json:"username"`
}
