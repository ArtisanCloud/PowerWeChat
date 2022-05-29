package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCorpVacationGetQuota struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
