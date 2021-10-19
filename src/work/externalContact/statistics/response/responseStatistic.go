package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseStatistic struct {
	*response.ResponseWork

	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`

	Items []*power.HashMap `json:"items"`
}
