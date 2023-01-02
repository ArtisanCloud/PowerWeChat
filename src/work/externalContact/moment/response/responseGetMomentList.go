package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetMomentList struct {
	response.ResponseWork
	NextCursor string         `json:"next_cursor"`
	MomentList *power.HashMap `json:"moment_list"`
}
