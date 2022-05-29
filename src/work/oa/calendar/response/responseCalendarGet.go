package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCalendarGet struct {
	*response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`
}
