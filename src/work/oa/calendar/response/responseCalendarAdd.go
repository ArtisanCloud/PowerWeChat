package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCalendarAdd struct {
	*response.ResponseWork

	CalID string `json:"cal_id"`

}
