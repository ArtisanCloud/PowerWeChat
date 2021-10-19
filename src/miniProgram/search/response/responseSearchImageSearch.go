package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseSearchImageSearch struct {
	*response.ResponseMiniProgram

	Items []*power.HashMap `json:"items"`
}
