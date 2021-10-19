package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseJournalGetStatList struct {
	*response.ResponseWork

	StatList *power.HashMap `json:"stat_list"`
}
