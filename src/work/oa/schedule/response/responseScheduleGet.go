package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseScheduleGet struct {
	*response.ResponseWork

	ScheduleList []*power.HashMap `json:"schedule_list"`
}
