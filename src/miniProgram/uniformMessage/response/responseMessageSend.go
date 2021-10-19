package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseDataCubeVisitInfo struct {
	*response.ResponseMiniProgram

	RefDate    string           `json:"ref_date"`
	VisitUVNew []*power.HashMap `json:"visit_uv_new"`
	VisitUV    []*power.HashMap `json:"visit_uv"`
}
