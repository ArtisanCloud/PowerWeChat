package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseIMGAICrop struct {
	response.ResponseMiniProgram
	Results []*power.HashMap `json:"results"`
}
