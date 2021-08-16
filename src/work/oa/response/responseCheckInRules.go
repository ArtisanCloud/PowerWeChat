package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCheckInRules struct {
	*response.ResponseWork
	Info []*object.HashMap `json:"info"`
}
