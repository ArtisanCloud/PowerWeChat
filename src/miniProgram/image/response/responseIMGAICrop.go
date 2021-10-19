package response

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
)

type ResponseIMGAICrop struct {
	*response.ResponseMiniProgram
	Results []*power.HashMap `json:"results"`
}
