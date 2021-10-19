package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCorpVacationGetQuota struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
