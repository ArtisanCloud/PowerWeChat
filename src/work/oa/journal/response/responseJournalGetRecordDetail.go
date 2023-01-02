package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseJournalGetRecordDetail struct {
	response.ResponseWork

	Info *power.HashMap `json:"info"`
}
