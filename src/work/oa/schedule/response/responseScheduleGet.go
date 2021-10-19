package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseScheduleGet struct {
	*response.ResponseWork

	ScheduleList []*power.HashMap `json:"schedule_list"`

}
