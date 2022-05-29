package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseBroadcastGetSubAnchor struct {
	*response.ResponseMiniProgram

	UserName string `json:"username"`
}
