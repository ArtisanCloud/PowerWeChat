package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseScheduleAdd struct {
	*response.ResponseWork

	ScheduleID string `json:"schedule_id"`

}
