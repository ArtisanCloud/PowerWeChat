package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpVacationGetConfig struct {
	*response.ResponseWork
	Lists []*object.HashMap `json:"lists"`
}

