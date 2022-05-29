package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type ResponseSearchImageSearch struct {
	*response.ResponseMiniProgram

	Items []*power.HashMap `json:"items"`
}
