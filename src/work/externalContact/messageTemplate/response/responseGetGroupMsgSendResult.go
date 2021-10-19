package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGetGroupMsgSendResult struct {
	*response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	SendList   []*power.HashMap `json:"send_list"`
}
