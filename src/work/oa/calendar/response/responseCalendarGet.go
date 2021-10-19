package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCalendarGet struct {
	*response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`

}
