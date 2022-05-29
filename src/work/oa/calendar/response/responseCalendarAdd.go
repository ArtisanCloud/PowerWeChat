package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCalendarAdd struct {
	*response.ResponseWork

	CalID string `json:"cal_id"`
}
