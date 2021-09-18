package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseListContactWay struct {
	*response.ResponseWork

	ContactWay []*power.HashMap `json:"contact_way"`
}
