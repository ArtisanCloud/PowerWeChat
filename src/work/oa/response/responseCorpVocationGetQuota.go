package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpVacationGetQuota struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
