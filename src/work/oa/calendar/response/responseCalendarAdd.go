package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCalendarAdd struct {
	*response.ResponseWork

	CalID string `json:"cal_id"`

}
