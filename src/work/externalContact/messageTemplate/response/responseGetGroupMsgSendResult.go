package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetGroupMsgSendResult struct {
	*response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	SendList   []*power.HashMap `json:"send_list"`
}
