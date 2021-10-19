package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	*response.ResponseMiniProgram

	List     []*power.HashMap `json:"list"`
}
