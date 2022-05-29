package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseScheduleAdd struct {
	*response.ResponseWork

	ScheduleID string `json:"schedule_id"`
}
