package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCheckInSchedulist struct {
	*response.ResponseWork

	ScheduleList []*power.HashMap `json:"schedule_list"`
}
