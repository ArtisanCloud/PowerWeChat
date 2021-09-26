package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetGroupMsgTask struct {
	*response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	TaskList   []*power.HashMap `json:"task_list"`
}
