package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseCorpVacationGetConfig struct {
	*response.ResponseWork
	Lists []*power.HashMap `json:"lists"`
}
