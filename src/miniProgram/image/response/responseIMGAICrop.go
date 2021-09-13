package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseIMGAICrop struct {
	*response.ResponseMiniProgram
	Results []*power.HashMap `json:"results"`
}
