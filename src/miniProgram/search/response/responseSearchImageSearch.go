package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseSearchImageSearch struct {
	*response.ResponseMiniProgram

	Items []*power.HashMap `json:"items"`
}
