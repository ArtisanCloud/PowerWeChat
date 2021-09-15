package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseGetMomentList struct {
	*response.ResponseWork
	NextCursor string         `json:"next_cursor"`
	MomentList *power.HashMap `json:"moment_list"`
}
