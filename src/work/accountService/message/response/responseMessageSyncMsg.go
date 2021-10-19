package response

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseMessageSyncMsg struct {
	*response.ResponseWork

	NextCursor string            `json:"next_cursor"`
	HasMore    int               `json:"has_more"`
	MsgList    []*object.HashMap `json:"msg_list"`
}
