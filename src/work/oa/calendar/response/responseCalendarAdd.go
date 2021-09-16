package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCalendarAdd struct {
	*response.ResponseWork

	CalID string `json:"cal_id"`

}
