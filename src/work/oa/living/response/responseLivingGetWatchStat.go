package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseLivingGetWatchStat struct {
	*response.ResponseWork

	NextKey   string   `json:"next_key"`
	StatInfo *power.HashMap `json:"stat_info"`
}
