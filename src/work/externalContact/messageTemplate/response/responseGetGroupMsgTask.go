package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseGetGroupMsgTask struct {
	*response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	TaskList   []*power.HashMap `json:"task_list"`
}
