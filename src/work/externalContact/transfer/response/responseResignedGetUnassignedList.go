package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseResignedGetUnassignedList struct {
	*response.ResponseWork

	Info       []*power.HashMap `json:"info"`
	IsLast     bool             `json:"is_last"`
	NextCursor string           `json:"next_cursor"`
}
