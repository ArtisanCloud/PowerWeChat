package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseCorpVacationGetQuota struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
