package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseMenuCreate struct {
	*response.ResponseWork
	Button []*power.HashMap `json:"button"`
}

