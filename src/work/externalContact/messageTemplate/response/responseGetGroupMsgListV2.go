package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetGroupMsgListV2 struct {
	*response.ResponseWork

	NextCursor   string           `json:"next_cursor"`
	GroupMsgList []*power.HashMap `json:"group_msg_list"`
}
