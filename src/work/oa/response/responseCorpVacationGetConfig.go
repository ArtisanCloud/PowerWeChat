package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseCorpVacationGetConfig struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
