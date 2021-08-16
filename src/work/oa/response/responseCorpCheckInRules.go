package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpCheckInRules struct {
	*response.ResponseWork
	Group []*object.HashMap `json:"group"`
}

