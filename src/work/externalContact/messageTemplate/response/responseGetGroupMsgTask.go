package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseGetGroupMsgTask struct {
	*response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	TaskList   []*power.HashMap `json:"task_list"`
}
