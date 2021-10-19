package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCalendarGet struct {
	*response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`

}
