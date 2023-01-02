package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type Task struct {
	UserID   string `json:"userid"`
	Status   int    `json:"status"`
	SendTime int    `json:"send_time"`
}

type ResponseGetGroupMsgTask struct {
	response.ResponseWork

	NextCursor string `json:"next_cursor"`

	TaskList []*Task `json:"task_list"`
}
