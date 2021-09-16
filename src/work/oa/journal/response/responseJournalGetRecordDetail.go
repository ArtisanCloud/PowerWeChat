package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseJournalGetRecordDetail struct {
	*response.ResponseWork

	Info *power.HashMap `json:"info"`

}
