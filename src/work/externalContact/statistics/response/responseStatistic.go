package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseStatistic struct {
	*response.ResponseWork

	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`

	Items []*power.HashMap `json:"items"`
}
