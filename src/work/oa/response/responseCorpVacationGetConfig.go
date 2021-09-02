package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpVacationGetConfig struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
