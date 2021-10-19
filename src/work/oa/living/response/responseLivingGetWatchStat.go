package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseLivingGetWatchStat struct {
	*response.ResponseWork

	NextKey   string   `json:"next_key"`
	StatInfo *power.HashMap `json:"stat_info"`
}
