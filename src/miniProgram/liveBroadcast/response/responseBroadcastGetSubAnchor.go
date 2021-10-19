package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	*response.ResponseMiniProgram

	UserName string `json:"username"`
}
