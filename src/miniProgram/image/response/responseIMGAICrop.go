package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseIMGAICrop struct {
	*response.ResponseMiniProgram
	Results []*power.HashMap `json:"results"`
}
