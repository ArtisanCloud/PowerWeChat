package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseJournalGetStatList struct {
	*response.ResponseWork

	StatList *power.HashMap `json:"stat_list"`
}
