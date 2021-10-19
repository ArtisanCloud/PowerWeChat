package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseDataCubeVisit struct {
	*response.ResponseMiniProgram

	RefDate string            `json:"ref_date"`
	List    []*power.HashMap `json:"list"`
}
